package utilities

import (
	"fmt"
	"reflect"
)

// Array.includes() for strings
func IncludesString(array []string, value string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

// Object.values()
func Values(object interface{}) []string {
	var list []string
	elements := reflect.ValueOf(&object).Elem()
	for i := 0; i < elements.NumField(); i++ {
		list = append(list, fmt.Sprintf("%v", elements.Field(i).Interface()))
	}
	return list
}
