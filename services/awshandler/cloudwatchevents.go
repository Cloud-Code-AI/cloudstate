package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

type cloudwatcheventsInfo struct {
	EventRules []types.Rule `json:"event_rules"`
}

// Gets all the Cloudwatch Event Data for a given regions and
// stores the results in output/{region}/cloudwatch/events.json file
func ListCloudwatchEventsFn(sdkConfig aws.Config) {
	const maxItems = 50

	// Create cloudwatch events service client
	client := cloudwatchevents.NewFromConfig(sdkConfig)
	cloudwatchData := cloudwatcheventsInfo{
		EventRules: getCloudwatchEventRules(client),
	}

	const (
		path = "/cloudwatch/events.json"
	)

	stats := addCloudwatchEventStats(cloudwatchData)
	output := BasicTemplate{
		Data:  cloudwatchData,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addCloudwatchEventStats(inp cloudwatcheventsInfo) interface{} {
	s := make(map[string]float64)
	s["event_rules"] = float64(len(inp.EventRules))
	return s
}

func getCloudwatchEventRules(client *cloudwatchevents.Client) []types.Rule {
	// Retrieve the tules
	result, err := client.ListRules(context.TODO(), &cloudwatchevents.ListRulesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var rules []types.Rule
	// TODO: Add pagination updates
	for _, dashboard := range result.Rules {

		rules = append(rules, dashboard)
	}
	return rules
}
