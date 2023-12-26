package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

// apigatewayData stores all the api info
type apigatewayData struct {
	Apis []types.RestApi `json:"apis"`
}

// Gets all the API Gateways for a given regions and
// stores the results in output/{region}/apigateway/data.json file
func apigatewayList(sdkConfig aws.Config) {
	// Create API Gateway service client
	client := apigateway.NewFromConfig(sdkConfig)

	// TODO: Add Pagination
	result, err := client.GetRestApis(context.TODO(), &apigateway.GetRestApisInput{})
	if err != nil {
		log.Printf("Couldn't find Apigateway Data. Here's why: %v\n", err)
	}

	data := apigatewayData{
		Apis: result.Items,
	}

	const (
		path = "/apigateway/data.json"
	)

	stats := addApigatewayStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing Apigateway Data")
	}

}

// Add stats for Apigateway Stats
func addApigatewayStats(inp apigatewayData) map[string]float64 {
	s := make(map[string]float64)
	s["stacks"] = float64(len(inp.Apis))
	return s
}
