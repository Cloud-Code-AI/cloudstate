package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

type elasticacheData struct {
	CacheClusters []types.CacheCluster `json:"cache_cluster"`
	Snapshots     []types.Snapshot     `json:"snapshots"`
}

// Gets all the elasticache Data for a given regions and
// stores the results in output/{region}/elasticache/data.json file
func ElasticCahceMetaData(sdkConfig aws.Config, parentpath string) {
	const maxItems = 50

	// Create elasticache service client
	client := elasticache.NewFromConfig(sdkConfig)
	data := elasticacheData{
		CacheClusters: getElasticacheClusters(client),
		Snapshots:     getElasticacheSnapshots(client),
	}

	const (
		path = "/elasticache/data.json"
	)

	stats := addElasticacheStats(data)
	output := BasicTemplate{
		Data:  data,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing elasticache data")
	}

}

func addElasticacheStats(inp elasticacheData) interface{} {
	s := make(map[string]float64)
	s["cache_clusters"] = float64(len(inp.CacheClusters))
	s["snapshots"] = float64(len(inp.Snapshots))
	return s
}

func getElasticacheClusters(client *elasticache.Client) []types.CacheCluster {
	// Retrieve the instances
	result, err := client.DescribeCacheClusters(context.TODO(), &elasticache.DescribeCacheClustersInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve clusters, %v", err)
	}
	var cache_clusters []types.CacheCluster
	// Process and print the instances details
	// TODO: Add pagination updates
	cache_clusters = result.CacheClusters
	return cache_clusters
}

func getElasticacheSnapshots(client *elasticache.Client) []types.Snapshot {
	result, err := client.DescribeSnapshots(context.TODO(), &elasticache.DescribeSnapshotsInput{})
	// TODO: Add pagination updates
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return result.Snapshots
}
