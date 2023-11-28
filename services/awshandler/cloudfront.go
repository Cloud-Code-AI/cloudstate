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

type CloudfrontList struct {
	Distributions []types.DistributionSummary `json:"websites"`
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

	output := CloudfrontList{
		Distributions: result.DistributionList.Items,
	}

	filepath := parentpath + sdkConfig.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing cloudfront distribution lists")
	}

}
