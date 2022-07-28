package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubernetes "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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

func main() {
	clientset := Connect2cluster()
	podlist, _ := clientset.CoreV1().Pods("kube-system").List(context.Background(), v1.ListOptions{LabelSelector: "k8s-app=metrics-server"})

	for _, j := range podlist.Items {
		fmt.Println(j.Name)
	}

}
