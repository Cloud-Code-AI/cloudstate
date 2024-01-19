package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Cloud-Code-AI/cloudstate/services/awshandler"
)

var (
	cloudProvider *string
	resourceType  *string
	region        *string
)

func main() {
	// Define commands
	gatherCmd := flag.NewFlagSet("gather", flag.ExitOnError)
	reportCmd := flag.NewFlagSet("report", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("Usage: cloudstate <command> [arguments]")
		return
	}

	switch os.Args[1] {
	case "gather":
		cloudProvider = gatherCmd.String("provider", "", "The cloud provider to interact with (e.g., 'aws', 'gcp', 'azure')")
		resourceType = gatherCmd.String("resource", "", "The type of resource to fetch (e.g., 'vm', 'storage', 'network')")
		region = gatherCmd.String("region", "", "The region for which the data should be fetched (e.g 'us-east-1', 'ap-south-1')")
		gatherCmd.Parse(os.Args[2:])
		gather()
	case "report":
		cloudProvider = reportCmd.String("provider", "", "The cloud provider to interact with (e.g., 'aws', 'gcp', 'azure')")
		reportCmd.Parse(os.Args[2:])
		generateReport()
	default:
		flag.Usage()
	}

}

func gather() {
	// Basic input validation
	if *cloudProvider == "" || *region == "" {
		fmt.Println("Missing required arguments for gather command")
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

func generateReport() {
	// Handle the cloud region based on the input
	switch *cloudProvider {
	case "aws":
		awshandler.GenerateAWSReport()
	case "gcp":
		fmt.Println("Report generation not implemented yet for gcp")
	case "azure":
		fmt.Println("Report generation not implemented yet for azure")
	default:
		fmt.Println("Unsupported cloud provider")
	}
}

func handleAWS(region, resourceType string) {
	// Implement AWS-specific logic here
	fmt.Printf("Provider: AWS \nregion: %s \n", region)
	awshandler.StoreAWSData(region)

}

func handleGCP(region, resourceType string) {
	// Implement GCP-specific logic here
	fmt.Printf("Provider: GCP \n region: %s on resource: %s\n", region, resourceType)

}

func handleAzure(region, resourceType string) {
	// Implement Azure-specific logic here
	fmt.Printf("Provider: Azure \n region: %s on resource: %s\n", region, resourceType)

}
