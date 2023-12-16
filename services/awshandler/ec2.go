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

type ec2Info struct {
	Instances []types.Instance `json:"Instances"`
	Snapshots []types.Snapshot `json:"Snapshots"`
	AMIs      []types.Image    `json:"AMIs"`
}

// Gets all the EC2 instance for a given regions and
// stores the results in output/{region}/ec2/instances.json file
func ListEc2Fn(sdkConfig aws.Config) {
	const maxItems = 50

	// Create Lambda service client
	client := ec2.NewFromConfig(sdkConfig)
	ec2Data := ec2Info{
		Instances: getEc2Instances(client),
		Snapshots: getEc2Snapshots(client),
		AMIs:      getEc2AMIs(client),
	}

	const (
		path = "/ec2/instances.json"
	)

	stats := addEc2tats(ec2Data)
	output := BasicTemplate{
		Data:  ec2Data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addEc2tats(inp ec2Info) interface{} {
	s := make(map[string]float64)
	s["instances"] = float64(len(inp.Instances))
	s["snapshots"] = float64(len(inp.Snapshots))
	s["amis"] = float64(len(inp.AMIs))
	return s
}

func getEc2Instances(client *ec2.Client) []types.Instance {
	// Retrieve the instances
	result, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var instances []types.Instance
	// Process and print the instances details
	// TODO: Add pagination updates
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}
	return instances
}

func getEc2Snapshots(client *ec2.Client) []types.Snapshot {
	result, err := client.DescribeSnapshots(context.TODO(), &ec2.DescribeSnapshotsInput{
		OwnerIds: []string{"self"},
	})
	// TODO: Add pagination updates
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return result.Snapshots
}

func getEc2AMIs(client *ec2.Client) []types.Image {
	result, err := client.DescribeImages(context.TODO(), &ec2.DescribeImagesInput{
		Owners: []string{"self"}, // Use "self" to list AMIs owned by your account

	})
	// TODO: Add pagination updates
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return result.Images
}
