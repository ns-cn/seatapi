package seatapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// SeaTableAction SeaTable操作
// 使用SeaTableAction执行SeaTable请求
// R: 返回数据
// E: 非200状态时解析到这个模型上
type SeaTableAction[R any] struct {
	Request  *http.Request  // 请求
	Response *http.Response // 响应
	Parsed   *R             // 成功之后的响应数据
	Data     []byte         // 数据
}

// DoWithDefaultClient 使用默认客户端执行请求
func (action *SeaTableAction[R]) DoWithDefaultClient() (*R, error) {
	return action.DoWithClient(http.DefaultClient)
}

// DoWithClient 使用自定义客户端执行请求
func (action *SeaTableAction[R]) DoWithClient(client *http.Client) (*R, error) {
	if action.Request == nil {
		return action.Parsed, fmt.Errorf("empty action")
	}
	var err error
	action.Response, err = client.Do(action.Request)
	if err != nil {
		return action.Parsed, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(action.Response.Body)
	if action.Response.StatusCode == http.StatusOK {
		if action.Response.Body != nil {
			action.Data, err = io.ReadAll(action.Response.Body)
			if err != nil {
				return nil, err
			}
			action.Parsed = new(R)
			err = json.Unmarshal(action.Data, action.Parsed)
			if err != nil {
				return action.Parsed, err
			}
		}
		return action.Parsed, nil
	} else if action.Response.Body != nil {
		action.Data, err = io.ReadAll(action.Response.Body)
	}
	return action.Parsed, fmt.Errorf(fmt.Sprintf("%d", action.Response.StatusCode))
}
