package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type lambdaList struct {
	Functions []types.FunctionConfiguration `json:"Functions"`
}

// Gets all the lambda functions for a given regions and
// stores the results in output/{region}/lambda/functions.json file
func ListLambdaFns(sdkConfig aws.Config) {
	const maxItems = 50

	// Create Lambda service client
	client := lambda.NewFromConfig(sdkConfig)

	var functions []types.FunctionConfiguration
	paginator := lambda.NewListFunctionsPaginator(client, &lambda.ListFunctionsInput{
		MaxItems: aws.Int32(int32(maxItems)),
	})
	for paginator.HasMorePages() && len(functions) < maxItems {
		pageOutput, err := paginator.NextPage(context.TODO())
		if err != nil {
			log.Panicf("Couldn't list functions for your account. Here's why: %v\n", err)
		}
		functions = append(functions, pageOutput.Functions...)
	}

	const (
		path = "/lambda/functions.json"
	)

	stats := addLambdaStats(functions)
	output := BasicTemplate{
		Data:  functions,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addLambdaStats(inp []types.FunctionConfiguration) interface{} {
	s := make(map[string]float64)
	s["instances"] = float64(len(inp))
	return s
}
