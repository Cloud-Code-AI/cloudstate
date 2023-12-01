package main

import (
	"flag"
	"fmt"

	"github.com/Cloud-Code-AI/cloudstate/services/awshandler"
)

func main() {
	// Define and parse command-line flags
	cloudProvider := flag.String("provider", "", "The cloud provider to interact with (e.g., 'aws', 'gcp', 'azure')")
	resourceType := flag.String("resource", "", "The type of resource to fetch (e.g., 'vm', 'storage', 'network')")
	region := flag.String("region", "", "The region for which the data should be fetched (e.g 'us-east-1', 'ap-south-1')")
	flag.Parse()

	// Basic input validation
	if *cloudProvider == "" || *region == "" {
		fmt.Println("Missing required arguments")
		flag.Usage()
		return
	}

	// Handle the cloud region based on the input
	switch *cloudProvider {
	case "aws":
		handleAWS(*region, *resourceType)
	case "gcp":
		handleGCP(*region, *resourceType)
	case "azure":
		handleAzure(*region, *resourceType)
	default:
		fmt.Println("Unsupported cloud provider")
	}
}

func handleAWS(region, resourceType string) {
	// Implement AWS-specific logic here
	fmt.Printf("Handling AWS region: %s \n", region)
	awshandler.StoreAWSData(region)

}

func handleGCP(region, resourceType string) {
	// Implement GCP-specific logic here
	fmt.Printf("Handling GCP region: %s on resource: %s\n", region, resourceType)

}

func handleAzure(region, resourceType string) {
	// Implement Azure-specific logic here
	fmt.Printf("Handling Azure region: %s on resource: %s\n", region, resourceType)

}
