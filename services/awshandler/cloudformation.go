package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

type CloudformationList struct {
	stacks []types.StackSummary `json:"stacks"`
}

// Gets all the distribution of cloudfront for a given regions and
// stores the results in output/{region}/cloudfront/distributions.json file
func CloudformationListFn(sdkConfig aws.Config) {
	// Create Cloudformation service client
	client := cloudformation.NewFromConfig(sdkConfig)

	result, err := client.ListStacks(context.TODO(), &cloudformation.ListStacksInput{})
	if err != nil {
		log.Printf("Couldn't list distribution. Here's why: %v\n", err)
	}

	const (
		path = "/cloudformation/stacks.json"
	)

	stats := addCloudformationStats(result.StackSummaries)
	output := BasicTemplate{
		Data:  result.StackSummaries,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing cloudfront distribution lists")
	}

}

// Add stats for cloudfront
func addCloudformationStats(inp []types.StackSummary) interface{} {
	s := make(map[string]float64)
	s["stacks"] = float64(len(inp))
	return s
}
