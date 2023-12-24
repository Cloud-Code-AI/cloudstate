package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBList struct {
	Tables []string `json:"Tables"`
}

// Gets all the dynamodb tables for a given regions and
// stores the results in output/{region}/dynamodb/tables.json file
func DynamoDBListFn(sdkConfig aws.Config) {
	// Create dynamodb service client
	client := dynamodb.NewFromConfig(sdkConfig)

	tables, err := client.ListTables(
		context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("Couldn't list tables. Here's why: %v\n", err)
	}

	const (
		path = "/dynamodb/tables.json"
	)

	stats := addDynamoDBStats(tables.TableNames)
	output := BasicTemplate{
		Data:  tables.TableNames,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addDynamoDBStats(inp []string) interface{} {
	s := make(map[string]float64)
	s["tables"] = float64(len(inp))
	return s
}
