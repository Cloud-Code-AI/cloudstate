package awshandler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

// A function to get All the AWS regions
func getAWSRegions() []string {
	return []string{
		"us-east-2",      // US East (Ohio)
		"us-east-1",      // US East (N. Virginia)
		"us-west-1",      // US West (N. California)
		"us-west-2",      // US West (Oregon)
		"af-south-1",     // Africa (Cape Town)
		"ap-east-1",      // Asia Pacific (Hong Kong)
		"ap-south-1",     // Asia Pacific (Mumbai)
		"ap-northeast-3", // Asia Pacific (Osaka)
		"ap-northeast-2", // Asia Pacific (Seoul)
		"ap-southeast-1", // Asia Pacific (Singapore)
		"ap-southeast-2", // Asia Pacific (Sydney)
		"ap-northeast-1", // Asia Pacific (Tokyo)
		"ca-central-1",   // Canada (Central)
		"cn-north-1",     // China (Beijing)
		"cn-northwest-1", // China (Ningxia)
		"eu-central-1",   // Europe (Frankfurt)
		"eu-west-1",      // Europe (Ireland)
		"eu-west-2",      // Europe (London)
		"eu-south-1",     // Europe (Milan)
		"eu-west-3",      // Europe (Paris)
		"eu-north-1",     // Europe (Stockholm)
		"me-south-1",     // Middle East (Bahrain)
		"sa-east-1",      // South America (São Paulo)
	}
}

// Creating a common interface for all the data points
type BasicTemplate struct {
	Stats interface{} `json:"stats"`
	Data  interface{} `json:"data"`
}

const parentpath = "output/"

func StoreAWSData(region string) {

	var regions []string

	// If the user wants to fetch all the region,
	// load them to regions variable
	if region == "all" {
		regions = getAWSRegions()
	} else {
		regions = append(regions, region)
	}

	for _, region := range regions {
		sdkConfig, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion(region),
		)
		if err != nil {
			fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
			fmt.Println(err)
			return
		}
		fmt.Printf("Storing data for region : %s\n", region)
		// Get all the S3 bucket data
		S3ListBucketss(sdkConfig)
		// Get all the lambda functions
		ListLambdaFns(sdkConfig)
		// Get all dynamodb tables
		DynamoDBListFn(sdkConfig)
		// Get EC2 Instance info
		ListEc2Fn(sdkConfig)
		// Get Cloudfront info
		CloudfrontListFn(sdkConfig)
		// Get IAM Infos
		IamMetadata(sdkConfig)
		// Get RDS Instance
		ListRDSFunc(sdkConfig)
		// Get Route53 Info
		ListRoute53Func(sdkConfig)
		// CloudFormation Stack
		CloudformationListFn(sdkConfig)
		// Cloudwatch Data
		ListCloudwatchFn(sdkConfig)
		apigatewayList(sdkConfig)
		ListCloudwatchEventsFn(sdkConfig)
		ListCloudwatchLogsFn(sdkConfig)
		KmsMetadata(sdkConfig)
		elasticsearchMetadata(sdkConfig)

	}

}
