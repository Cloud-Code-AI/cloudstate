package awshandler

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Cloud-Code-AI/cloudstate/services/utils"
)

func GenerateAWSReport() {
	// Get the most recent data stored for a provider
	dir, err := utils.GetMostRecentDirectory("./output/aws/")
	if err != nil {
		fmt.Println("Failed to find the resource meta data locally. Make sure to run gather command before running repor!")
		fmt.Println(err)
		return
	}

	// Compiles and list all the stats in a single file.
	allStats := make(map[string]interface{})
	var serviceName string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".json" {
			data, err := utils.ReadJSONFile(path)
			if err != nil {
				return err
			}
			serviceName = strings.Split(path, "/")[3]

			// Print the entire data for debugging
			fmt.Printf("Service name %v, data is %v, \n", serviceName, data["stats"])

			// Type assert 'stats' as map[string]int
			stats, ok := data["stats"].(map[string]interface{})
			if !ok {
				// handle the error: data["stats"] is not of type map[string]int
				fmt.Printf("Service name %v, data is %v, \n", serviceName, data["stats"])
				fmt.Printf("Error in parsing stats data for service %v\n", serviceName)
			} else {
				allStats[serviceName] = stats
			}
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	utils.PrintNested(allStats, "", 0)

	err = utils.WriteJSONToFile("output/aws/report.json", allStats)
	if err != nil {
		fmt.Println("Failed to Write the report file to json")
	}

}
