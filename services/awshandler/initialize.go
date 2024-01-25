package awshandler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

// A function to get All the AWS regions
func getAWSRegions() []string {
	// Currently supports US Only
	return []string{
		"us-east-2", // US East (Ohio)
		"us-east-1", // US East (N. Virginia)
		"us-west-1", // US West (N. California)
		"us-west-2", // US West (Oregon)
		// "af-south-1",     // Africa (Cape Town)
		// "ap-east-1",      // Asia Pacific (Hong Kong)
		// "ap-south-1",     // Asia Pacific (Mumbai)
		// "ap-northeast-3", // Asia Pacific (Osaka)
		// "ap-northeast-2", // Asia Pacific (Seoul)
		// "ap-southeast-1", // Asia Pacific (Singapore)
		// "ap-southeast-2", // Asia Pacific (Sydney)
		// "ap-northeast-1", // Asia Pacific (Tokyo)
		// "ca-central-1",   // Canada (Central)
		// "cn-north-1",     // China (Beijing)
		// "cn-northwest-1", // China (Ningxia)
		// "eu-central-1",   // Europe (Frankfurt)
		// "eu-west-1",      // Europe (Ireland)
		// "eu-west-2",      // Europe (London)
		// "eu-south-1",     // Europe (Milan)
		// "eu-west-3",      // Europe (Paris)
		// "eu-north-1",     // Europe (Stockholm)
		// "me-south-1",     // Middle East (Bahrain)
		// "sa-east-1",      // South America (São Paulo)
	}
}

// Creating a common interface for all the data points
type BasicTemplate struct {
	Stats interface{} `json:"stats"`
	Data  interface{} `json:"data"`
}

func StoreAWSData(region string) {

	var regions []string

	// If the user wants to fetch all the region,
	// load them to regions variable
	if region == "all" {
		regions = getAWSRegions()
	} else {
		regions = append(regions, region)
	}

	// parentpath := "output/aws/" + time.Now().Format("2006-01-02T15:04:05") + "/"
	parentpath := "output/aws/"

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

		// Run this only once as they are global resource
		if region == "us-east-1" {
			S3ListBucketss(sdkConfig, parentpath)
			ListRoute53Func(sdkConfig, parentpath)
		}
		// Get all the lambda functions
		ListLambdaFns(sdkConfig, parentpath)
		// Get all dynamodb tables
		DynamoDBListFn(sdkConfig, parentpath)
		// Get EC2 Instance info
		ListEc2Fn(sdkConfig, parentpath)
		// Get Cloudfront info
		CloudfrontListFn(sdkConfig, parentpath)
		// Get IAM Infos
		IamMetadata(sdkConfig, parentpath)
		// Get RDS Instance
		ListRDSFunc(sdkConfig, parentpath)
		// CloudFormation Stack
		CloudformationListFn(sdkConfig, parentpath)
		// Cloudwatch Data
		ListCloudwatchFn(sdkConfig, parentpath)
		apigatewayList(sdkConfig, parentpath)
		ListCloudwatchEventsFn(sdkConfig, parentpath)
		ListCloudwatchLogsFn(sdkConfig, parentpath)
		KmsMetadata(sdkConfig, parentpath)
		elasticsearchMetadata(sdkConfig, parentpath)
		ElasticCahceMetaData(sdkConfig, parentpath)
		ECRMetaData(sdkConfig, parentpath)
		codebuildMetadata(sdkConfig, parentpath)

	}

}
