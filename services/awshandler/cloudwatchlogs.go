package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

type cloudwatchlogsData struct {
	Logs []types.LogGroup `json:"logs"`
}

// Gets all the Cloudwatch Logs Data for a given regions and
// stores the results in output/{region}/cloudwatch/logs.json file
func ListCloudwatchLogsFn(sdkConfig aws.Config) {
	const maxItems = 50

	// Create cloudwatch events service client
	client := cloudwatchlogs.NewFromConfig(sdkConfig)
	data := cloudwatchlogsData{
		Logs: getCloudwatchLogGroups(client),
	}

	const (
		path = "/cloudwatch/logs.json"
	)

	stats := addCloudwatchLogStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addCloudwatchLogStats(inp cloudwatchlogsData) interface{} {
	s := make(map[string]float64)
	s["logs"] = float64(len(inp.Logs))
	return s
}

func getCloudwatchLogGroups(client *cloudwatchlogs.Client) []types.LogGroup {
	// Retrieve the tules
	result, err := client.DescribeLogGroups(context.TODO(), &cloudwatchlogs.DescribeLogGroupsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var logGroups []types.LogGroup
	// TODO: Add pagination updates
	for _, logGroup := range result.LogGroups {

		logGroups = append(logGroups, logGroup)
	}
	return logGroups
}
