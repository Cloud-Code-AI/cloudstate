package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func PrintNested(data interface{}, prefix string, level int) {
	indent := "    " // 4 spaces
	if level == 0 {
		// Assuming top-level is map[string]map[string]interface{}
		for region, services := range data.(map[string]map[string]interface{}) {
			fmt.Println(prefix + strings.Repeat(indent, level) + region)
			PrintNested(services, prefix, level+1)
		}
	} else if reflect.TypeOf(data).Kind() == reflect.Map {
		// Handle nested maps
		for k, v := range data.(map[string]interface{}) {
			valueType := reflect.TypeOf(v).Kind()
			keyString := fmt.Sprintf("%-30s", prefix+strings.Repeat(indent, level)+k)
			if valueType == reflect.Map || valueType == reflect.Slice {
				fmt.Println(keyString)
				PrintNested(v, prefix, level+1)
			} else {
				fmt.Printf("%s %-20v \n", keyString, v)
			}
		}
	} else if reflect.TypeOf(data).Kind() == reflect.Slice {
		// Handle slices
		for _, v := range data.([]interface{}) {
			valueType := reflect.TypeOf(v).Kind()
			keyString := fmt.Sprintf("%-30s", prefix+strings.Repeat(indent, level))
			if valueType == reflect.Map || valueType == reflect.Slice {
				fmt.Println(keyString)
				PrintNested(v, prefix, level+1)
			} else {
				fmt.Printf("%s %-20v \n", keyString, v)
			}
		}
	}
}
