package utils

import (
	"encoding/json"
	"reflect"
)

func ConvertStructData(object interface{}, targetStruct interface{}) interface{} {
	// Use reflection to get the type of the targetted struct
	targetStructType := reflect.TypeOf(targetStruct)
	// Create a new variable of the same type as the targetted struct
	targetStructData := reflect.New(targetStructType).Interface()
	// Convert object to fit into the targetted struct
	temporaryVariable, _ := json.Marshal(object)
	json.Unmarshal(temporaryVariable, &targetStructData)

	return targetStructData
}
