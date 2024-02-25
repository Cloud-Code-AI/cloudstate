package gcphandler

// A function to get All the AWS regions
func getGoogleRegions() []string {
	// Currently supports US Only
	return []string{
		"us-central1", // US East (Ohio)
	}
}

// Creating a common interface for all the data points
type BasicTemplate struct {
	Stats interface{} `json:"stats"`
	Data  interface{} `json:"data"`
}

func StoreGoogleData(region string, outFolder string) {

	var regions []string

	// If the user wants to fetch all the region,
	// load them to regions variable
	if region == "all" {
		regions = getGoogleRegions()
	} else {
		regions = append(regions, region)
	}

	projectID := "test"

	// parentpath := "output/aws/" + time.Now().Format("2006-01-02T15:04:05") + "/"
	if outFolder == "" {
		outFolder = "output"
	}
	parentpath := outFolder + "/gcp/"

	for _, region := range regions {

		// Run this only once as they are global resource
		getComputeInfo(projectID, parentpath, region)

	}

}
