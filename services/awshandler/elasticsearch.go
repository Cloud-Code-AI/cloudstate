package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
)

// elasticsearchData holds information about Elastic Search Data.
type elasticsearchData struct {
	Domains []types.DomainInfo         `json:"domains"`
	VPCs    []types.VpcEndpointSummary `json:"vpcs"`
}

// Gets all the Elastic Search Data for a given regions and
// stores the results in output/{region}/ess/keys.json file
func elasticsearchMetadata(sdkConfig aws.Config, parentpath string) {
	// Create Elasticsearch service client
	client := elasticsearchservice.NewFromConfig(sdkConfig)

	const (
		path = "/ess/data.json"
	)

	data := elasticsearchData{
		Domains: getESSDomains(client),
		VPCs:    getESSVPC(client),
	}
	stats := addESSStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing Elastic Search Service data to json file")
	}

}

// Add stats for Elastic search
func addESSStats(info elasticsearchData) interface{} {
	stats := make(map[string]float64)
	stats["domains"] = float64(len(info.Domains))
	return stats
}

func getESSDomains(client *elasticsearchservice.Client) []types.DomainInfo {
	var domains []types.DomainInfo
	// TODO: Add Pagination to the list users
	result, err := client.ListDomainNames(
		context.TODO(),
		&elasticsearchservice.ListDomainNamesInput{},
	)
	domains = result.DomainNames
	if err != nil {
		log.Printf("Couldn't list domains. Here's why: %v\n", err)
	}
	return domains
}

func getESSVPC(client *elasticsearchservice.Client) []types.VpcEndpointSummary {
	var vpcs []types.VpcEndpointSummary
	// TODO: Add Pagination to the list users
	result, err := client.ListVpcEndpoints(
		context.TODO(),
		&elasticsearchservice.ListVpcEndpointsInput{},
	)
	vpcs = result.VpcEndpointSummaryList
	if err != nil {
		log.Printf("Couldn't list domains. Here's why: %v\n", err)
	} else {

		for result.NextToken != nil {
			result, err = client.ListVpcEndpoints(
				context.TODO(),
				&elasticsearchservice.ListVpcEndpointsInput{},
			)
			if err != nil {
				log.Printf("Couldn't list domains. Here's why: %v\n", err)
				break
			}
			vpcs = append(vpcs, result.VpcEndpointSummaryList...)
		}
	}
	return vpcs
}
