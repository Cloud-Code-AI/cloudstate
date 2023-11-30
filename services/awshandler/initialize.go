package awshandler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
)

func StoreAWSData(region string) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
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

}
