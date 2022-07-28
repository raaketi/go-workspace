package main

import (
	"context"
	"log"
	"rajasureshaditya/go-workspace/K8sapis/getk8scalls"
)

type aditya getk8scalls.Object

func main() {
	myclientset := *getk8scalls.Connect2cluster()
	// getk8scalls.LaunchK8sJob(&myclientset)
	log.Println(getk8scalls.GetDeployments(&myclientset, context.Background(), "default"))
}
