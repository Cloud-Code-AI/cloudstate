package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
)

// Cloudfront data stores the distribution summaries
type cloudfrontData struct {
	Distributions []types.DistributionSummary `json:"distributions"`
}

// Gets all the distribution of cloudfront for a given regions and
// stores the results in output/{region}/cloudfront/distributions.json file
func CloudfrontListFn(sdkConfig aws.Config) {
	// Create cloudfront service client
	client := cloudfront.NewFromConfig(sdkConfig)

	result, err := client.ListDistributions(context.TODO(), &cloudfront.ListDistributionsInput{})
	if err != nil {
		log.Printf("Couldn't list distribution. Here's why: %v\n", err)
	}

	const (
		path = "/cloudfront/distributions.json"
	)

	data := cloudfrontData{
		Distributions: result.DistributionList.Items,
	}
	stats := addCloudfrontStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing cloudfront distribution lists")
	}
}

// Add stats for cloudfront
func addCloudfrontStats(inp cloudfrontData) map[string]float64 {
	s := make(map[string]float64)
	s["websites"] = float64(len(inp.Distributions))
	return s
}
