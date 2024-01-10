package awshandler

import (
	"context"
	"fmt"
	"log"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
)

// TODO: For the images, we can list them better and just have names of the images.

// codebuildData holds information about Codebuild Projects, Builds, and enviorments.
type codebuildData struct {
	Projects         []string                    `json:"projects"`
	Builds           []string                    `json:"builds"`
	EnviormentImages []types.EnvironmentPlatform `json:"enviorment_images"`
}

// Gets all the Cloud Build Data for a given regions and
// stores the results in output/{region}/codebuild/data.json file
func codebuildMetadata(sdkConfig aws.Config, parentpath string) {
	// Create CodeBuild service client
	client := codebuild.NewFromConfig(sdkConfig)

	const (
		path = "/codebuild/data.json"
	)

	codebuildResult := codebuildData{
		Projects:         listCodeBuildProjects(client),
		Builds:           listCodeBuilds(client),
		EnviormentImages: listCodeBuildEnviorments(client),
	}
	stats := addCodeBuildStats(codebuildResult)
	output := BasicTemplate{
		Data:  codebuildResult,
		Stats: stats,
	}

	filepath := parentpath + sdkConfig.Region + path

	err := utils.WriteJSONToFile(filepath, output)
	if err != nil {
		fmt.Println("Error writing codebuild data")
	}

}

// Add stats for codebuild
func addCodeBuildStats(info codebuildData) interface{} {
	stats := make(map[string]float64)
	stats["projects"] = float64(len(info.Projects))
	stats["builds"] = float64(len(info.Builds))
	stats["enviroment_images"] = float64(len(info.EnviormentImages))
	return stats
}

func listCodeBuildProjects(client *codebuild.Client) []string {
	var projects []string
	result, err := client.ListProjects(context.TODO(), &codebuild.ListProjectsInput{})
	if err != nil {
		log.Printf("Couldn't list codebuild projects. Here's why: %v\n", err)
	} else {
		projects = result.Projects
		for result.NextToken != nil {
			result, err = client.ListProjects(context.TODO(), &codebuild.ListProjectsInput{
				NextToken: result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list codebuild projects. Here's why: %v\n", err)
				break
			}
			projects = append(projects, result.Projects...)
		}
	}
	return projects
}

func listCodeBuilds(client *codebuild.Client) []string {
	var builds []string
	result, err := client.ListBuilds(context.TODO(), &codebuild.ListBuildsInput{})
	if err != nil {
		log.Printf("Couldn't list codebuild builds. Here's why: %v\n", err)
	} else {
		builds = result.Ids
		for result.NextToken != nil {
			result, err = client.ListBuilds(context.TODO(), &codebuild.ListBuildsInput{
				NextToken: result.NextToken,
			})
			if err != nil {
				log.Printf("Couldn't list codebuild builds. Here's why: %v\n", err)
				break
			}
			builds = append(builds, result.Ids...)
		}
	}
	return builds
}

func listCodeBuildEnviorments(client *codebuild.Client) []types.EnvironmentPlatform {
	var enviorments []types.EnvironmentPlatform
	result, err := client.ListCuratedEnvironmentImages(
		context.TODO(), &codebuild.ListCuratedEnvironmentImagesInput{})
	if err != nil {
		log.Printf("Couldn't list codebuild Images. Here's why: %v\n", err)
	} else {
		enviorments = result.Platforms
	}
	return enviorments
}
