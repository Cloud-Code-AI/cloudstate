package gcphandler

import (
	"context"
	"fmt"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"google.golang.org/api/iterator"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

type ec2Info struct {
	Instances []*computepb.Instance `json:"Instances"`
}

// Gets all the EC2 instance for a given regions and
// stores the results in output/{region}/ec2/instances.json file
func getComputeInfo(projectID string, parentpath string, region string) {
	const maxItems = 50

	// Create EC2 service client
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		fmt.Println("NewInstancesRESTClient:")
	}
	defer instancesClient.Close()

	ec2Data := ec2Info{
		Instances: getComputeInstances(projectID, region, ctx, instancesClient),
	}

	const (
		path = "/compute/instances.json"
	)

	stats := addEc2stats(ec2Data)
	output := BasicTemplate{
		Data:  ec2Data,
		Stats: stats,
	}

	filepath := parentpath + region + path

	err = utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing lambda function lists")
	}

}

func addEc2stats(inp ec2Info) interface{} {
	s := make(map[string]float64)
	s["instances"] = float64(len(inp.Instances))
	return s
}

func getComputeInstances(projectID string, zone string, ctx context.Context, instancesClient *compute.InstancesClient) []*computepb.Instance {
	req := &computepb.ListInstancesRequest{
		Zone: zone,
	}
	var instances []*computepb.Instance
	it := instancesClient.List(ctx, req)
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error fetching instance: ")
			fmt.Println(err)
			break
		}
		instances = append(instances, instance)
	}
	return instances
}
