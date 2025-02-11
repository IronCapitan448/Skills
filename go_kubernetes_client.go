package main

import (
	"context" // Provides context management for API requests
	"fmt"     // Used for printing output to console
	"log"     // Used for logging errors and info messages
	"os"      // Used to interact with environment variables and file paths

	// Kubernetes client libraries
	"k8s.io/client-go/kubernetes"               // Main client library to interact with Kubernetes
	"k8s.io/client-go/tools/clientcmd"          // Helps load the kubeconfig file
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // Used for defining metadata like namespace, labels, etc.
)

func main() {
	// Define the path to the kubeconfig file (usually found in ~/.kube/config)
	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	// Load the kubeconfig file to create a client configuration
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err) // Fatal log if kubeconfig cannot be loaded
	}

	// Create a Kubernetes client using the loaded configuration
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err) // Fatal log if client cannot be created
	}

	// Define the namespace to query (default namespace)
	namespace := "default"

	// List all pods in the specified namespace
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err) // Fatal log if pods cannot be listed
	}

	// Print the names of all retrieved pods
	fmt.Println("Pods in the 'default' namespace:")
	for _, pod := range pods.Items {
		fmt.Println("- ", pod.Name)
	}
}
