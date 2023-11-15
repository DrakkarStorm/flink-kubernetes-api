package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Message struct pour stocker le message de réponse JSON
type FlinkDeployment struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// Get operator tenants list
func ListFlinkDeployment() ([]FlinkDeployment, error) {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	dynamic := dynamic.NewForConfigOrDie(config)

	resourceId := schema.GroupVersionResource{
		Group:    "flink.apache.org",
		Version:  "v1beta1",
		Resource: "flinkdeployments",
	}

	list, err := dynamic.Resource(resourceId).Namespace("flink-operator").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var deployments []FlinkDeployment
	for _, item := range list.Items {
		name, _, err := unstructured.NestedString(item.UnstructuredContent(), "metadata", "name")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		status, _, err := unstructured.NestedString(item.UnstructuredContent(), "metadata", "state")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		cpu, _, err := unstructured.NestedString(item.UnstructuredContent(), "metadata", "clusterInfo", "total-cpu")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		memory, _, err := unstructured.NestedString(item.UnstructuredContent(), "metadata", "clusterInfo", "total-memory")
		if err != nil {
			log.Println(err)
			return nil, err
		}

		deployments = append(deployments, FlinkDeployment{Name: name, Status: status, CPU: cpu, Memory: memory})
	}

	return deployments, nil
}
