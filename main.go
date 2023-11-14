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

func main() {
	ListFlinkDeployment()
	log.Println("End")
}

// Get operator tenants list
func ListFlinkDeployment() ([]string, error) {
	log.Println("Begin")

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
		log.Print(err)
		return nil, err
	}

	log.Print(len(list.Items))

	var tenantList []string
	for _, item := range list.Items {

		name, _, err := unstructured.NestedString(item.UnstructuredContent(), "metadata", "name")
		if err != nil {
			log.Print(err)

			return nil, err
		}
		log.Println(name)
	}
	return tenantList, nil
}
