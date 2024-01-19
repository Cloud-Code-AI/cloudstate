package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func PrintNested(data interface{}, prefix string, level int) {
	indent := "    " // 4 spaces
	if reflect.TypeOf(data).Kind() == reflect.Map {
		for k, v := range data.(map[string]interface{}) {
			valueType := reflect.TypeOf(v).Kind()
			keyString := fmt.Sprintf("%-30s", prefix+strings.Repeat(indent, level)+k)
			if valueType == reflect.Map || valueType == reflect.Slice {
				fmt.Println(keyString, "(nested)")
				PrintNested(v, prefix, level+1)
			} else {
				fmt.Println(keyString, fmt.Sprintf("%-20v", v))
			}
		}
	} else if reflect.TypeOf(data).Kind() == reflect.Slice {
		for i, v := range data.([]interface{}) {
			valueType := reflect.TypeOf(v).Kind()
			keyString := fmt.Sprintf("%-30s", prefix+strings.Repeat(indent, level)+fmt.Sprintf("[%d]", i))
			if valueType == reflect.Map || valueType == reflect.Slice {
				fmt.Println(keyString, "(nested)")
				PrintNested(v, prefix, level+1)
			} else {
				fmt.Println(keyString, fmt.Sprintf("%-20v", v))
			}
		}
	}
}
