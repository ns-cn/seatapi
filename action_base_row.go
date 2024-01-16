package seatapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (api SeaTableApi) ListRowsWithSQL(ctx BaseContext, sql string, convertKeys bool) SeaTableAction[RowsWithSQL[map[string]interface{}]] {
	url := "https://cloud.seatable.io/dtable-db/api/v1/query/68fa7023-aa53-467d-b000-b2edbe844a34/"
	payload := strings.NewReader(fmt.Sprintf("{\"convert_keys\":%t,\"sql\":\"%s\"}", convertKeys, sql))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", ctx.AccessToken))
	action := SeaTableAction[RowsWithSQL[map[string]interface{}]]{}
	action.Request = req
	return action
}

// RowsWithSQL 结构体
// Result Result的数据类型
type RowsWithSQL[Result any] struct {
	ApiResult
	Metadata []struct {
		Key  string `json:"key"`
		Name string `json:"name"`
		Type string `json:"type"`
		Data *struct {
			Options                []interface{} `json:"options"`
			DefaultValue           interface{}   `json:"default_value"`
			EnableFillDefaultValue bool          `json:"enable_fill_default_value,omitempty"`
		} `json:"data"`
	} `json:"metadata"`
	Results []Result `json:"results"`
}

// ParseRowsWithSQL 将 RowsWithSQL 结构体解析为 RowsWithSQL 结构体(强制转换Result结构体)
func ParseRowsWithSQL[F any, T any](from RowsWithSQL[F], to *RowsWithSQL[T]) {
	to.ApiResult = from.ApiResult
	to.Metadata = from.Metadata
	bytes, _ := json.Marshal(from.Results)
	var data T
	_ = json.Unmarshal(bytes, &data)
	to.Results = append(to.Results, data)
}
