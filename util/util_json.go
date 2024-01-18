package util

import (
	"encoding/json"
	"reflect"
)

// ParseToJsonString 将data转换为jsonString
func ParseToJsonString(data interface{}) string {
	if data == nil {
		return ""
	} else {
		// 将data转换为jsonString
		bytes, err := json.Marshal(data)
		if err != nil {
			return ""
		}
		return string(bytes)
	}
}

// ParseFromMapToPointer 将from转换为to
// toPointer: 必须标注为指针类型
func ParseFromMapToPointer(from interface{}, toPointer interface{}) {
	// 判断to必须是指针类型
	if reflect.TypeOf(toPointer).Kind() != reflect.Ptr {
		panic("to must be a pointer")
	}
	if from != nil {
		if bytes, err := json.Marshal(from); err == nil {
			_ = json.Unmarshal(bytes, toPointer)
		}
	}
}
