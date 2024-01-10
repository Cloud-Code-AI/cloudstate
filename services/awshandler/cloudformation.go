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

// cloudformationData stores all the stack info
type cloudformationData struct {
	Stacks []types.StackSummary `json:"stacks"`
}

// Gets all the stacks of cloudformation for a given regions and
// stores the results in output/{region}/cloudfromation/stacks.json file
func CloudformationListFn(sdkConfig aws.Config, parentpath string) {
	// Create Cloudformation service client
	client := cloudformation.NewFromConfig(sdkConfig)

	result, err := client.ListStacks(context.TODO(), &cloudformation.ListStacksInput{})
	if err != nil {
		log.Printf("Couldn't list  stacks of cloudformation. Here's why: %v\n", err)
	}

	data := cloudformationData{
		Stacks: result.StackSummaries,
	}

	const (
		path = "/cloudformation/stacks.json"
	)

	stats := addCloudformationStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing stacks of cloudformation lists")
	}

}

// Add stats for cloudformation
func addCloudformationStats(inp cloudformationData) interface{} {
	s := make(map[string]float64)
	s["stacks"] = float64(len(inp.Stacks))
	return s
}
