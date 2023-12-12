package utils

import "fmt"

func AuditStat(region string) {
	// for a given region, get basic audit stats and
	// store them for dashboard publishing
	filePaths := GetJSONFiles(region)

	for _, filepath := range filePaths {
		jsonData, err := ReadJSONFile(filepath)
		if err != nil {
			fmt.Println("Error reading json file: ", filepath)
		}
		fmt.Println(jsonData)
		// TODO: Read the json files and merge them to create a specific json content
	}

}
