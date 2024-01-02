package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
)

type ecrData struct {
	Repositories []types.Repository `json:"repositories"`
}

// Gets all the ecr Data for a given regions and
// stores the results in output/{region}/ecr/data.json file
func ECRMetaData(sdkConfig aws.Config) {
	const maxItems = 50

	// Create ecr service client
	client := ecr.NewFromConfig(sdkConfig)
	data := ecrData{
		Repositories: getECRRepository(client),
	}

	const (
		path = "/ecr/data.json"
	)

	stats := addECRStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing ecr data")
	}

}

func addECRStats(inp ecrData) interface{} {
	s := make(map[string]float64)
	s["repositories"] = float64(len(inp.Repositories))
	return s
}

func getECRRepository(client *ecr.Client) []types.Repository {
	var repositories []types.Repository

	result, err := client.DescribeRepositories(context.TODO(), &ecr.DescribeRepositoriesInput{MaxResults: aws.Int32(100)})
	if err != nil {
		log.Printf("Couldn't list ecr repos. Here's why: %v\n", err)
	} else {
		repositories = result.Repositories
		for result.NextToken != nil {
			result, err = client.DescribeRepositories(context.TODO(), &ecr.DescribeRepositoriesInput{
				MaxResults: aws.Int32(100),
				NextToken:  result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't ecr repos. Here's why: %v\n", err)
				break
			}
			repositories = append(repositories, result.Repositories...)
		}
	}
	return repositories
}
