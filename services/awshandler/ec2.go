package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type ec2List struct {
	Instances []types.Instance `json:"Instances"`
}

// Gets all the EC2 instance for a given regions and
// stores the results in output/{region}/ec2/instances.json file
func ListEc2Fn(sdkConfig aws.Config) {
	const maxItems = 50

	// Create Lambda service client
	client := ec2.NewFromConfig(sdkConfig)

	// Retrieve the instances
	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var instances []types.Instance
	// Process and print the instances details
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}

	const (
		path = "/ec2/instances.json"
	)

	output := ec2List{
		Instances: instances,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}
