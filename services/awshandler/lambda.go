package awshandler

import (
	"fmt"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type lambdaList struct {
	Functions []*lambda.FunctionConfiguration `json:"Buckets"`
}

// Gets all the lambda functions for a given regions and
// stores the results in output/{region}/lambda/functions.json file
func ListLambdaFns(sess *session.Session) {
	// Create Lambda service client
	svc := lambda.New(sess)

	result, err := svc.ListFunctions(nil)
	if err != nil {
		utils.ExitErrorf("Unable to list functions, %v", err)
	}

	const (
		path = "/lambda/functions.json"
	)

	output := lambdaList{
		Functions: result.Functions,
	}

	filepath := parentpath + *sess.Config.Region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}
