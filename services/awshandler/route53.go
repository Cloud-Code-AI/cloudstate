package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
)

type Route53Info struct {
	Domains []types.DomainSummary `json:"Domains"`
}

// Gets all the Route53 Domains for a given regions and
// stores the results in output/{region}/route53/instances.json file
func ListRoute53Func(sdkConfig aws.Config, parentpath string) {
	const maxItems = 50

	// Create Route53 service client
	client := route53domains.NewFromConfig(sdkConfig)
	route53Data := Route53Info{
		Domains: getRoute53Domains(client),
	}

	const (
		path = "/route53/domains.json"
	)

	stats := addRoute53Stats(route53Data)
	output := BasicTemplate{
		Data:  route53Data,
		Stats: stats,
	}

	filepath := parentpath + "global" + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addRoute53Stats(inp Route53Info) interface{} {
	s := make(map[string]float64)
	s["domains"] = float64(len(inp.Domains))
	return s
}

func getRoute53Domains(client *route53domains.Client) []types.DomainSummary {
	// Retrieve the instances
	result, err := client.ListDomains(context.TODO(), &route53domains.ListDomainsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var domains []types.DomainSummary
	// Process and print the instances details
	// TODO: Add pagination updates
	for _, instance := range result.Domains {
		domains = append(domains, instance)
	}
	return domains
}
