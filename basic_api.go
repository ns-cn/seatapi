package seatapi

import (
	"fmt"
)

// SeaTableApi  接口
type SeaTableApi struct {
	Host string
}

// wholeUrl 生成完整URL
func (api SeaTableApi) wholeUrl(path string) string {
	return fmt.Sprintf("%s%s", api.Host, path)
}

// ApiError 仅错误时返回
type ApiError struct {
	ErrorMsg string `json:"error_msg"`
}

// ApiResult API操作结果,可能成功也可能失败
type ApiResult struct {
	ErrorMessage string `json:"error_message"`
	Success      bool   `json:"success"`
}
