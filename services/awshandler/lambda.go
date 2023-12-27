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
	Functions []types.FunctionConfiguration `json:"functions"`
	Layers    []types.LayersListItem        `json:"layers"`
}

// Gets all the lambda functions for a given regions and
// stores the results in output/{region}/lambda/functions.json file
func ListLambdaFns(sdkConfig aws.Config) {

	// Create Lambda service client
	client := lambda.NewFromConfig(sdkConfig)

	data := lambdaList{
		Layers:    listLambdaLayers(client),
		Functions: listFunctions(client),
	}

	const (
		path = "/lambda/functions.json"
	)

	stats := addLambdaStats(data)
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

func addLambdaStats(inp lambdaList) interface{} {
	s := make(map[string]float64)
	s["functions"] = float64(len(inp.Functions))
	s["layers"] = float64(len(inp.Layers))
	return s
}

func listFunctions(client *lambda.Client) []types.FunctionConfiguration {
	var functions []types.FunctionConfiguration
	result, err := client.ListFunctions(context.TODO(),
		&lambda.ListFunctionsInput{
			MaxItems: aws.Int32(int32(50)),
		})
	if err != nil {
		log.Printf("Couldn't list lambda functions. Here's why: %v\n", err)
	} else {
		functions = result.Functions
		for result.NextMarker != nil {
			result, err = client.ListFunctions(
				context.TODO(),
				&lambda.ListFunctionsInput{
					MaxItems: aws.Int32(50),
					Marker:   result.NextMarker,
				},
			)
			if err != nil {
				log.Printf("Couldn't list lambda functions. Here's why: %v\n", err)
				break
			}
			functions = append(functions, result.Functions...)
		}
	}
	return functions
}

func listLambdaLayers(client *lambda.Client) []types.LayersListItem {
	var layers []types.LayersListItem
	result, err := client.ListLayers(context.TODO(),
		&lambda.ListLayersInput{
			MaxItems: aws.Int32(50),
		},
	)
	if err != nil {
		log.Printf("Couldn't list lambda layers. Here's why: %v\n", err)
	} else {
		layers = result.Layers
		for result.NextMarker != nil {
			result, err = client.ListLayers(context.TODO(), &lambda.ListLayersInput{
				MaxItems: aws.Int32(50),
				Marker:   result.NextMarker,
			})
			if err != nil {
				log.Printf("Couldn't list lambda layers. Here's why: %v\n", err)
				break
			}
			layers = append(layers, result.Layers...)
		}
	}
	return layers
}
