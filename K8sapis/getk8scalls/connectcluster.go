package getk8scalls

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	depV1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	clientcmd "k8s.io/client-go/tools/clientcmd"
)

func GetDeployments(clientset *kubernetes.Clientset, ctx context.Context, namespace string, v1 metav1.ListOptions) ([]depV1.Deployment, error) {
	list, err := clientset.AppsV1().Deployments(namespace).
		List(ctx, v1)
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func Connect2cluster() *kubernetes.Clientset {
	home, exists := os.LookupEnv("HOME")
	if !exists {
		home = "/root"
	}

	configPath := filepath.Join(home, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		log.Panicln("failed to create K8s config")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panicln("Failed to create K8s clientset")
	}

	return clientset
}

func LaunchK8sJob(clientset *kubernetes.Clientset) {
	jobs := clientset.BatchV1().Jobs("default")
	var backOffLimit int32 = 0
	var cmd string = "ls -aRil"

	jobSpec := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "raja",
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "raja",
							Image:   "alpine:latest",
							Command: strings.Split(cmd, " "),
						},
					},
					RestartPolicy: v1.RestartPolicyNever,
				},
			},
			BackoffLimit: &backOffLimit,
		},
	}

	_, err := jobs.Create(context.TODO(), jobSpec, metav1.CreateOptions{})
	if err != nil {
		log.Fatalln("Failed to create K8s job.", err)
	}

	//print job details
	log.Println("Created K8s job successfully")
}
