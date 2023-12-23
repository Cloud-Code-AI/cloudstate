package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

type IAMInfos struct {
	Users    []types.User   `json:"Users"`
	Policies []types.Policy `json:"Policies"`
	Roles    []types.Role   `json:"Roles"`
}

// Gets all the IAM Data for a given regions and
// stores the results in output/{region}/IAM/iam.json file
func IAMList(sdkConfig aws.Config) {
	// Create cloudfront service client
	IamClient := iam.NewFromConfig(sdkConfig)

	userList := getIAMUsers(IamClient)
	policyList := listPolicies(IamClient)
	rolesList := listRoles(IamClient)
	const (
		path = "/IAM/iam.json"
	)

	IamResult := IAMInfos{
		Users:    userList,
		Policies: policyList,
		Roles:    rolesList,
	}
	stats := addIAMStats(IamResult)
	output := BasicTemplate{
		Data:  IamResult,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing cloudfront distribution lists")
	}

}

// Add stats for cloudfront
func addIAMStats(info IAMInfos) interface{} {
	s := make(map[string]float64)
	s["users"] = float64(len(info.Users))
	s["roles"] = float64(len(info.Roles))
	s["policies"] = float64(len(info.Policies))
	return s
}

func getIAMUsers(IamClient *iam.Client) []types.User {
	var users []types.User
	// TODO: Add Pagination to the list users
	result, err := IamClient.ListUsers(context.TODO(), &iam.ListUsersInput{
		MaxItems: aws.Int32(100),
	})
	if err != nil {
		log.Printf("Couldn't list users. Here's why: %v\n", err)
	} else {

		users = result.Users
	}
	return users
}

func listPolicies(IamClient *iam.Client) []types.Policy {
	var policies []types.Policy
	// TODO: Add Pagination to the list policies
	result, err := IamClient.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
		MaxItems: aws.Int32(100),
	})
	if err != nil {
		log.Printf("Couldn't list policies. Here's why: %v\n", err)
	} else {
		policies = result.Policies
		for result.IsTruncated {
			result, err = IamClient.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
				MaxItems: aws.Int32(100),
				Marker:   result.Marker,
			})
			if err != nil {
				log.Printf("Couldn't list policies. Here's why: %v\n", err)
				break
			}
			policies = append(policies, result.Policies...)

		}
	}
	return policies
}

// ListRoles gets up to maxRoles roles.
func listRoles(IamClient *iam.Client) []types.Role {
	var roles []types.Role
	result, err := IamClient.ListRoles(context.TODO(),
		&iam.ListRolesInput{MaxItems: aws.Int32(100)},
	)
	if err != nil {
		log.Printf("Couldn't list roles. Here's why: %v\n", err)
	} else {
		roles = result.Roles
	}
	return roles
}
