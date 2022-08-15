package main

import (
	"context"
	"fmt"
	"log"
	"rajasureshaditya/go-workspace/K8sapis/getk8scalls"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	myclientset := getk8scalls.Connect2cluster()
	// getk8scalls.LaunchK8sJob(&myclientset)
	log.Println(getk8scalls.GetDeployments(myclientset, context.Background(), "default", metav1.ListOptions{LabelSelector: "app=football"}))
	footDeployments, _ := getk8scalls.GetDeployments(myclientset, context.Background(), "default", metav1.ListOptions{LabelSelector: "app=football"})
	for _, i := range footDeployments {
		fmt.Println(i.Name)
	}
}
