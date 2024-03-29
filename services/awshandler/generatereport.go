package awshandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
)

func GenerateAWSReport(outFolder string) {
	// Get the most recent data stored for a provider

	if outFolder == "" {
		outFolder = "output"
	}
	dir := outFolder + "/aws/"

	// Compiles and list all the stats in a single file.
	regionStats := make(map[string]map[string]interface{})
	allData := make(map[string]map[string]interface{})
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := utils.ReadJSONFile(path)
			if err != nil {
				return err
			}
			// Extract service name and region from path
			pathParts := strings.Split(path, "/")
			if len(pathParts) < 4 {
				return fmt.Errorf("unexpected path format: %s", path)
			}
			path_len := len(pathParts)
			serviceName := pathParts[path_len-2]
			regionName := pathParts[path_len-3]

			// Print the entire data for debugging
			fmt.Printf("Service name %v, data is %v, \n", serviceName, data["stats"])

			// Type assert 'stats' as map[string]int
			stats, ok := data["stats"].(map[string]interface{})
			if !ok {
				return fmt.Errorf("invalid stats format for service: %s, region: %s", serviceName, regionName)
			}

			// Group stats by region
			if _, exists := regionStats[regionName]; !exists {
				regionStats[regionName] = make(map[string]interface{})
				allData[regionName] = make(map[string]interface{})
			}
			regionStats[regionName][serviceName] = stats
			allData[regionName][serviceName] = data
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print and write the report
	utils.PrintNested(regionStats, "", 0)

	// Storing report data
	err = utils.WriteJSONToFile(outFolder+"/aws_report.json", regionStats)
	if err != nil {
		fmt.Println("Failed to Write the report file to json")
		fmt.Println(err)
	}

	// Storing report data
	err = utils.WriteJSONToFile(outFolder+"/aws_metadata.json", allData)
	if err != nil {
		fmt.Println("Failed to Write the all data file to json")
		fmt.Println(err)
	}

}
