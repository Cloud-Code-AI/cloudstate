package awshandler

import (
	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func initSess() *session.Session {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	if err != nil {
		utils.ExitErrorf("Unable to create session, %v", err)
	}

	return sess
}

func StoreAWSData() {
	sess := initSess()
	// Get all the S3 bucket data
	S3ListBucketss(sess)
	// Get all the lambda functions
	ListLambdaFns(sess)

}
