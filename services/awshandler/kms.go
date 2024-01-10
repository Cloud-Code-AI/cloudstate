package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// kmsData holds information about KMS keys.
type kmsData struct {
	Keys []types.KeyListEntry `json:"keys"`
}

// Gets all the KMS Data for a given regions and
// stores the results in output/{region}/KMS/keys.json file
func KmsMetadata(sdkConfig aws.Config, parentpath string) {
	// Create KMS service client
	client := kms.NewFromConfig(sdkConfig)

	const (
		path = "/kms/keys.json"
	)

	data := kmsData{
		Keys: getKMSKeys(client),
	}
	stats := addKMSStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing kms data to json file")
	}

}

// Add stats for KMS
func addKMSStats(info kmsData) interface{} {
	stats := make(map[string]float64)
	stats["keys"] = float64(len(info.Keys))
	return stats
}

func getKMSKeys(client *kms.Client) []types.KeyListEntry {
	var keys []types.KeyListEntry
	// TODO: Add Pagination to the list users
	result, err := client.ListKeys(context.TODO(), &kms.ListKeysInput{
		Limit: aws.Int32(1000),
	})
	if err != nil {
		log.Printf("Couldn't list keys. Here's why: %v\n", err)
	} else {
		keys = result.Keys
		for result.Truncated {
			result, err = client.ListKeys(context.TODO(), &kms.ListKeysInput{
				Limit:  aws.Int32(100),
				Marker: result.NextMarker,
			})
			if err != nil {
				log.Printf("Couldn't list keys. Here's why: %v\n", err)
				break
			}
			keys = append(keys, result.Keys...)
		}
	}
	return keys
}
