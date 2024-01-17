package util

import "encoding/json"

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

// ParseFromMapToAny 将from转换为to
// to: 必须标注为指针类型
func ParseFromMapToAny(from interface{}, to interface{}) {
	if from != nil {
		if bytes, err := json.Marshal(from); err != nil {
			_ = json.Unmarshal(bytes, to)
		}
	}
}
