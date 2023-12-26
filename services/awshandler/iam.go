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

// IAMData holds information about IAM users, policies, and roles.
type IAMData struct {
	Users    []types.User   `json:"Users"`
	Policies []types.Policy `json:"Policies"`
	Roles    []types.Role   `json:"Roles"`
}

// Gets all the IAM Data for a given regions and
// stores the results in output/{region}/IAM/iam.json file
func IAMList(sdkConfig aws.Config) {
	// Create IAM service client
	iamClient := iam.NewFromConfig(sdkConfig)

	userList := getIAMUsers(iamClient)
	policyList := listPolicies(iamClient)
	rolesList := listRoles(iamClient)
	const (
		path = "/IAM/iam.json"
	)

	IamResult := IAMData{
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
func addIAMStats(info IAMData) interface{} {
	stats := make(map[string]float64)
	stats["users"] = float64(len(info.Users))
	stats["roles"] = float64(len(info.Roles))
	stats["policies"] = float64(len(info.Policies))
	return stats
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
		for result.IsTruncated {
			result, err = IamClient.ListUsers(context.TODO(), &iam.ListUsersInput{
				MaxItems: aws.Int32(100),
				Marker:   result.Marker,
			})
			if err != nil {
				log.Printf("Couldn't list policies. Here's why: %v\n", err)
				break
			}
			users = append(users, result.Users...)
		}
	}
	return users
}

// List IAM policies which are created by users
func listPolicies(IamClient *iam.Client) []types.Policy {
	var policies []types.Policy
	result, err := IamClient.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
		MaxItems: aws.Int32(100),
		Scope:    "Local",
	})
	if err != nil {
		log.Printf("Couldn't list policies. Here's why: %v\n", err)
	} else {
		policies = result.Policies
		for result.IsTruncated {
			result, err = IamClient.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
				MaxItems: aws.Int32(100),
				Marker:   result.Marker,
				Scope:    "Local",
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
		&iam.ListRolesInput{
			MaxItems: aws.Int32(100),
		},
	)
	if err != nil {
		log.Printf("Couldn't list roles. Here's why: %v\n", err)
	} else {
		roles = result.Roles
		for result.IsTruncated {
			result, err = IamClient.ListRoles(context.TODO(), &iam.ListRolesInput{
				MaxItems: aws.Int32(100),
				Marker:   result.Marker,
			})
			if err != nil {
				log.Printf("Couldn't list policies. Here's why: %v\n", err)
				break
			}
			roles = append(roles, result.Roles...)
		}
	}
	return roles
}
