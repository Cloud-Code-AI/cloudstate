package awshandler

import (
	"context"
	"fmt"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type s3Buckets struct {
	Buckets []types.Bucket `json:"Buckets"`
}

const (
	jsonpath   = "/s3/buckets.json"
	parentpath = "output/"
)

// Gets all the files from s3 for a given regions and
// stores the results in output/s3/buckets.json file
func S3ListBucketss(sdkConfig aws.Config) {
	// Create S3 service client
	client := s3.NewFromConfig(sdkConfig)

	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		utils.ExitErrorf("Unable to list buckets, %v", err)
	}

	// fmt.Println(result.Buckets)

	stats := addS3Stats(result.Buckets)
	output := BasicTemplate{
		Data:  result.Buckets,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + jsonpath

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing S3 bucket lists")
	}

}

func addS3Stats(inp []types.Bucket) interface{} {
	s := make(map[string]float64)
	s["buckets"] = float64(len(inp))
	return s
}
