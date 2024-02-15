package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
)

type autoscalingInfo struct {
	AutoScalingGroups    []types.AutoScalingGroup           `json:"groups"`
	AutoScalingInstances []types.AutoScalingInstanceDetails `json:"instances"`
	LaunchConfigurations []types.LaunchConfiguration        `json:"metric_streams"`
}

// Gets all the AutoScaling Data for a given regions and
// stores the results in output/{region}/autoscaling/metrics.json file
func ListAutoscalingFn(sdkConfig aws.Config, parentpath string) {
	const maxItems = 50

	// Create autoscaling service client
	client := autoscaling.NewFromConfig(sdkConfig)
	data := autoscalingInfo{
		AutoScalingGroups:    getAutoScalingGroups(client),
		AutoScalingInstances: getAutoScalingInstances(client),
		LaunchConfigurations: getLaunchConfigurations(client),
	}

	const (
		path = "/autoscaling/data.json"
	)

	stats := addAutoscalingStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing autoscaling data")
	}

}

func addAutoscalingStats(inp autoscalingInfo) interface{} {
	s := make(map[string]float64)
	s["autoscaling_groups"] = float64(len(inp.AutoScalingGroups))
	s["autoscaling_instances"] = float64(len(inp.AutoScalingInstances))
	s["launch_configurations"] = float64(len(inp.LaunchConfigurations))
	return s
}

func getAutoScalingGroups(client *autoscaling.Client) []types.AutoScalingGroup {
	// Retrieve the AutoScaling dashboard
	result, err := client.DescribeAutoScalingGroups(context.TODO(), &autoscaling.DescribeAutoScalingGroupsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve autoscaling groups, %v", err)
	}
	var groups []types.AutoScalingGroup
	// TODO: Add pagination updates
	groups = result.AutoScalingGroups
	return groups
}

func getAutoScalingInstances(client *autoscaling.Client) []types.AutoScalingInstanceDetails {
	// Retrieve the AutoScaling Instances
	result, err := client.DescribeAutoScalingInstances(context.TODO(), &autoscaling.DescribeAutoScalingInstancesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve autoscaling instances, %v", err)
	}
	var instances []types.AutoScalingInstanceDetails
	// TODO: Add pagination updates
	instances = result.AutoScalingInstances
	return instances
}

func getLaunchConfigurations(client *autoscaling.Client) []types.LaunchConfiguration {
	// Retrieve the AutoScaling LaunchConfigurations
	result, err := client.DescribeLaunchConfigurations(context.TODO(), &autoscaling.DescribeLaunchConfigurationsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve autoscaling configs, %v", err)
	}
	var configs []types.LaunchConfiguration
	// TODO: Add pagination updates
	configs = result.LaunchConfigurations
	return configs
}
