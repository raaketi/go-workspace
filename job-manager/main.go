package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// Configure log to include timestamps
	log.SetFlags(log.LstdFlags)

	// Get environment variables
	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		log.Fatalf("NAMESPACE environment variable is not set")
	}
	jobName := os.Getenv("JOB_NAME")
	if jobName == "" {
		log.Fatalf("JOB_NAME environment variable is not set")
	}
	cronJobName := os.Getenv("CRONJOB_NAME")
	if cronJobName == "" {
		log.Fatalf("CRONJOB_NAME environment variable is not set")
	}

	// Read configurable wait time from environment variable (default to 20 minutes if not set)
	waitTimeStr := os.Getenv("WAIT_TIME")
	waitTime := 20 // default wait time in minutes
	if waitTimeStr != "" {
		var err error
		waitTime, err = strconv.Atoi(waitTimeStr)
		if err != nil {
			log.Printf("Invalid WAIT_TIME value '%s', using default 20 minutes.\n", waitTimeStr)
			waitTime = 20
		}
	}

	// Load in-cluster Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to load in-cluster configuration: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	// Repeat the entire process every configurable wait time
	for {
		log.Println("Starting job management cycle...")

		// Step 1: Continuously check for the CronJob every 5 minutes until it's available
		for {
			log.Printf("Checking for cronjob %s...\n", cronJobName)
			err = checkCronJobExists(clientset, namespace, cronJobName)
			if err != nil {
				log.Printf("Cronjob %s not found. Retrying in 5 minutes...\n", cronJobName)
				time.Sleep(5 * time.Minute) // Retry after 5 minutes if cronjob is not found
				continue
			}

			log.Printf("CronJob %s found.\n", cronJobName)
			break // Exit loop when cronjob is found
		}

		// Step 2: Proceed with Job management after CronJob is confirmed to exist
		for {
			// Check if the job exists and is completed
			jobCompleted := checkAndDeleteJob(clientset, namespace, jobName)

			// If the job was completed (or not found), create a new job from the cronjob
			if jobCompleted {
				// Wait for 10 seconds after deletion before creating a new job
				log.Println("Waiting for 10 seconds to ensure the job is fully deleted...")
				time.Sleep(10 * time.Second) // Wait for 10 seconds

				err = createJobFromCronJob(clientset, namespace, cronJobName, jobName)
				if err != nil {
					log.Printf("Failed to create job from cronjob %s: %v\n", cronJobName, err)
				} else {
					log.Printf("Job %s created successfully from cronjob %s.\n", jobName, cronJobName)
				}
				break // Exit the job management loop after successful job creation
			}

			// Add a short sleep to avoid tight loops
			time.Sleep(5 * time.Second)
		}

		// Step 3: Wait for the configurable time before repeating the cycle
		log.Printf("Waiting for %d minutes before the next cycle...\n", waitTime)
		time.Sleep(time.Duration(waitTime) * time.Minute) // Use configurable wait time
	}
}

// checkCronJobExists checks if the specified cronjob exists in the namespace
func checkCronJobExists(clientset *kubernetes.Clientset, namespace, cronJobName string) error {
	cronJobsClient := clientset.BatchV1().CronJobs(namespace)

	_, err := cronJobsClient.Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		return logErrorf("CronJob %s not found: %v", cronJobName, err)
	}

	log.Printf("CronJob %s is available.\n", cronJobName)
	return nil
}

// checkAndDeleteJob checks if the job exists and is completed, then deletes it if necessary
// with PropagationPolicy set to "Orphan" to leave the pods as orphaned.
func checkAndDeleteJob(clientset *kubernetes.Clientset, namespace, jobName string) bool {
	jobsClient := clientset.BatchV1().Jobs(namespace)

	job, err := jobsClient.Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		log.Printf("Job %s not found or error occurred: %v\n", jobName, err)
		return true // If the job doesn't exist, proceed to creation
	}

	// Check job status
	for _, condition := range job.Status.Conditions {
		if condition.Type == v1.JobComplete && condition.Status == "True" {
			log.Printf("Job %s is completed. Deleting the job but leaving pods as orphaned...\n", jobName)
			deletePolicy := metav1.DeletePropagationOrphan
			err := jobsClient.Delete(context.TODO(), jobName, metav1.DeleteOptions{
				PropagationPolicy: &deletePolicy,
			})
			if err != nil {
				log.Fatalf("Failed to delete job: %v", err)
			}
			log.Printf("Job %s deletion initiated, leaving pods as orphaned.\n", jobName)
			return true // Job was completed and deleted
		}
	}

	log.Printf("Job %s is not completed or still running.\n", jobName)
	return false
}

// createJobFromCronJob creates a new job from a cronjob
func createJobFromCronJob(clientset *kubernetes.Clientset, namespace, cronJobName, newJobName string) error {
	cronJobsClient := clientset.BatchV1().CronJobs(namespace)

	// Get the cronjob
	cronJob, err := cronJobsClient.Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		return logErrorf("Failed to get cronjob %s: %v", cronJobName, err)
	}

	// Create a job from the cronjob's spec
	job := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newJobName,
			Namespace: namespace,
		},
		Spec: *cronJob.Spec.JobTemplate.Spec.DeepCopy(), // Copy the job spec from cronjob
	}

	jobsClient := clientset.BatchV1().Jobs(namespace)
	_, err = jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		return logErrorf("Failed to create job %s: %v", newJobName, err)
	}

	return nil
}

// logErrorf is a helper function to log errors and return an error object
func logErrorf(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	log.Println(err)
	return err
}
