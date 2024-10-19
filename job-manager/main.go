package main

import (
	"context"
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
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

	// Load in-cluster Kubernetes configuration
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Failed to load in-cluster configuration: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create clientset: %v", err)
	}

	// Step 1: Delete the job if it is completed
	jobCompleted := checkAndDeleteJob(clientset, namespace, jobName)

	// Step 2: Create a new job from the cronjob if the previous job was deleted
	if jobCompleted {
		createJobFromCronJob(clientset, namespace, cronJobName, jobName)
	}
}

func checkAndDeleteJob(clientset *kubernetes.Clientset, namespace, jobName string) bool {
	jobsClient := clientset.BatchV1().Jobs(namespace)

	job, err := jobsClient.Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		fmt.Printf("Job %s not found or error occurred: %v\n", jobName, err)
		return true // Proceed to create a new job if it doesn't exist
	}

	for _, condition := range job.Status.Conditions {
		if condition.Type == v1.JobComplete && condition.Status == "True" {
			fmt.Printf("Job %s is completed. Deleting...\n", jobName)
			err := jobsClient.Delete(context.TODO(), jobName, metav1.DeleteOptions{})
			if err != nil {
				log.Fatalf("Failed to delete job: %v", err)
			}
			fmt.Printf("Job %s deleted successfully.\n", jobName)
			return true // Job was completed and deleted
		}
	}

	fmt.Printf("Job %s is not completed or still running.\n", jobName)
	return false
}

func createJobFromCronJob(clientset *kubernetes.Clientset, namespace, cronJobName, newJobName string) {
	cronJobsClient := clientset.BatchV1().CronJobs(namespace)

	cronJob, err := cronJobsClient.Get(context.TODO(), cronJobName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get cronjob %s: %v", cronJobName, err)
	}

	job := &v1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newJobName,
			Namespace: namespace,
		},
		Spec: *cronJob.Spec.JobTemplate.Spec.DeepCopy(),
	}

	jobsClient := clientset.BatchV1().Jobs(namespace)
	_, err = jobsClient.Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create job %s: %v", newJobName, err)
	}

	fmt.Printf("Job %s created successfully from cronjob %s.\n", newJobName, cronJobName)
}
