package awshandler

import (
	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func InitSess() *session.Session {
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

func StoreData(services []string) {
	// for _, service := range services {
	// 	fetchService()
	// }
}
