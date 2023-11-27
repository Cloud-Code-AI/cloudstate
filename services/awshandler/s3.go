package awshandler

import (
	"fmt"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Buckets struct {
	Buckets []*s3.Bucket `json:"Buckets"`
}

const (
	jsonpath   = "/s3/buckets.json"
	parentpath = "output/"
)

// Gets all the files from s3 for a given regions and
// stores the results in output/s3/buckets.json file
func S3ListBucketss(sess *session.Session) {
	// Create S3 service client
	svc := s3.New(sess)

	result, err := svc.ListBuckets(nil)
	if err != nil {
		utils.ExitErrorf("Unable to list buckets, %v", err)
	}

	output := s3Buckets{
		Buckets: result.Buckets,
	}

	filepath := parentpath + *sess.Config.Region + jsonpath

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing S3 bucket lists")
	}

}
