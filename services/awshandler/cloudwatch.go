package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

type cloudwatchInfo struct {
	Dashboards    []types.DashboardEntry    `json:"Dashboards"`
	Metrics       []types.Metric            `json:"Metrics"`
	MetricStreams []types.MetricStreamEntry `json:"MetricStreams"`
}

// Gets all the Cloudwatch Data for a given regions and
// stores the results in output/{region}/cloudwatch/metrics.json file
func ListCloudwatchFn(sdkConfig aws.Config) {
	const maxItems = 50

	// Create cloudwatch service client
	client := cloudwatch.NewFromConfig(sdkConfig)
	cloudwatchData := cloudwatchInfo{
		Dashboards:    getCloudwatchDashboards(client),
		Metrics:       getCloudwatchMetrics(client),
		MetricStreams: getCloudwatchMetricStreams(client),
	}

	const (
		path = "/cloudwatch/metrics.json"
	)

	stats := addCloudwatchStats(cloudwatchData)
	output := BasicTemplate{
		Data:  cloudwatchData,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addCloudwatchStats(inp cloudwatchInfo) interface{} {
	s := make(map[string]float64)
	s["dashboards"] = float64(len(inp.Dashboards))
	s["metrics"] = float64(len(inp.Metrics))
	s["metric_streams"] = float64(len(inp.MetricStreams))
	return s
}

func getCloudwatchDashboards(client *cloudwatch.Client) []types.DashboardEntry {
	// Retrieve the instances
	result, err := client.ListDashboards(context.TODO(), &cloudwatch.ListDashboardsInput{})
	if err != nil {
		log.Fatalf("Unable to retrieve instances, %v", err)
	}
	var dashboards []types.DashboardEntry
	// Process and print the instances details
	// TODO: Add pagination updates
	for _, dashboard := range result.DashboardEntries {

		dashboards = append(dashboards, dashboard)
	}
	return dashboards
}

func getCloudwatchMetrics(client *cloudwatch.Client) []types.Metric {
	result, err := client.ListMetrics(context.TODO(), &cloudwatch.ListMetricsInput{})
	// TODO: Add pagination updates
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return result.Metrics
}

func getCloudwatchMetricStreams(client *cloudwatch.Client) []types.MetricStreamEntry {
	result, err := client.ListMetricStreams(context.TODO(), &cloudwatch.ListMetricStreamsInput{})
	// TODO: Add pagination updates
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return result.Entries
}
