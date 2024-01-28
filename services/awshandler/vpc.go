package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

type vpcInfo struct {
	VpcEndpoints []types.VpcEndpoint `json:"vpc_endpoints"`
	Vpcs         []types.Vpc         `json:"vpcs"`
	Subnets      []types.Subnet      `json:"subnets"`
	RouteTables  []types.RouteTable  `json:"route_table"`
}

// Gets all the EC2 instance for a given regions and
// stores the results in output/{region}/ec2/instances.json file
func ListVpcFn(sdkConfig aws.Config, parentpath string) {
	const maxItems = 50

	// Create EC2 service client
	client := ec2.NewFromConfig(sdkConfig)
	vpcData := vpcInfo{
		VpcEndpoints: getVpcEndpoints(client),
		Vpcs:         getVpcs(client),
		Subnets:      getSubnets(client),
		RouteTables:  getRouteTable(client),
	}

	const (
		path = "/vpc/data.json"
	)

	stats := addVPCStats(vpcData)
	output := BasicTemplate{
		Data:  vpcData,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addVPCStats(i vpcInfo) interface{} {
	s := make(map[string]float64)
	s["vpcs"] = float64(len(i.Vpcs))
	s["vpc_endpoints"] = float64(len(i.VpcEndpoints))
	s["subnets"] = float64(len(i.Subnets))
	s["route_tables"] = float64(len(i.RouteTables))
	return s
}

func getVpcEndpoints(client *ec2.Client) []types.VpcEndpoint {
	var vpc_endpoints []types.VpcEndpoint
	result, err := client.DescribeVpcEndpoints(context.TODO(), &ec2.DescribeVpcEndpointsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}

	if err != nil {
		log.Printf("Couldn't list VPC endpoints. Here's why: %v\n", err)
	} else {
		vpc_endpoints = result.VpcEndpoints
		for result.NextToken != nil {
			result, err = client.DescribeVpcEndpoints(context.TODO(), &ec2.DescribeVpcEndpointsInput{
				MaxResults: aws.Int32(1000),
				NextToken:  result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list VPC endpoints. Here's why: %v\n", err)
				break
			}
			vpc_endpoints = append(vpc_endpoints, result.VpcEndpoints...)
		}
	}
	return vpc_endpoints
}

func getVpcs(client *ec2.Client) []types.Vpc {
	var vpcs []types.Vpc
	result, err := client.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{})
	if err != nil {
		log.Fatalf("Unable to list VPCs, %v", err)
	}

	if err != nil {
		log.Printf("Couldn't list VPCs. Here's why: %v\n", err)
	} else {
		vpcs = result.Vpcs
		for result.NextToken != nil {
			result, err = client.DescribeVpcs(context.TODO(), &ec2.DescribeVpcsInput{
				MaxResults: aws.Int32(1000),
				NextToken:  result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list VPCs. Here's why: %v\n", err)
				break
			}
			vpcs = append(vpcs, result.Vpcs...)
		}
	}
	return vpcs
}

func getSubnets(client *ec2.Client) []types.Subnet {
	var subnets []types.Subnet
	result, err := client.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve Subnets, %v", err)
	}

	if err != nil {
		log.Printf("Couldn't list VPC Subnets. Here's why: %v\n", err)
	} else {
		subnets = result.Subnets
		for result.NextToken != nil {
			result, err = client.DescribeSubnets(context.TODO(), &ec2.DescribeSubnetsInput{
				MaxResults: aws.Int32(1000),
				NextToken:  result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list VPC Subnets. Here's why: %v\n", err)
				break
			}
			subnets = append(subnets, result.Subnets...)
		}
	}
	return subnets
}

func getRouteTable(client *ec2.Client) []types.RouteTable {
	var route_tables []types.RouteTable
	result, err := client.DescribeRouteTables(context.TODO(), &ec2.DescribeRouteTablesInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve Subnets, %v", err)
	}

	if err != nil {
		log.Printf("Couldn't list VPC Subnets. Here's why: %v\n", err)
	} else {
		route_tables = result.RouteTables
		for result.NextToken != nil {
			result, err = client.DescribeRouteTables(context.TODO(), &ec2.DescribeRouteTablesInput{
				MaxResults: aws.Int32(1000),
				NextToken:  result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list VPC Subnets. Here's why: %v\n", err)
				break
			}
			route_tables = append(route_tables, result.RouteTables...)
		}
	}
	return route_tables
}
