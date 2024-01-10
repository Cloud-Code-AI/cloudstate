package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

type RDSInfo struct {
	Instances []types.DBInstance `json:"Instances"`
	// Snapshots []types.Snapshot `json:"Snapshots"`
	// AMIs      []types.Image    `json:"AMIs"`
}

// Gets all the rds instance for a given regions and
// stores the results in output/{region}/ec2/instances.json file
func ListRDSFunc(sdkConfig aws.Config, parentpath string) {
	const maxItems = 50

	// Create RDS service client
	client := rds.NewFromConfig(sdkConfig)
	rdsData := RDSInfo{
		Instances: getRDSInstances(client),
	}

	const (
		path = "/rds/instances.json"
	)

	stats := addRDSStats(rdsData)
	output := BasicTemplate{
		Data:  rdsData,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addRDSStats(inp RDSInfo) interface{} {
	s := make(map[string]float64)
	s["instances"] = float64(len(inp.Instances))
	return s
}

func getRDSInstances(client *rds.Client) []types.DBInstance {
	// Retrieve the instances
	result, err := client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var instances []types.DBInstance
	// Process and print the instances details
	// TODO: Add pagination updates
	for _, instance := range result.DBInstances {
		instances = append(instances, instance)
	}
	return instances
}
