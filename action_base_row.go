package seatapi

import (
	"fmt"
	"github.com/ns-cn/seatapi/util"
	"net/http"
	"strings"
)

// ListRowsWithSQL 使用SQL方式获取行数据
// convertKeys 确定列是作为其键 （true） 返回，还是作为其名称（默认为 false）返回。
func (api SeaTableApi) ListRowsWithSQL(ctx BaseContext, sql string, convertKeys bool) SeaTableAction[RowsWithSQL[map[string]interface{}]] {
	url := api.wholeUrl("/dtable-db/api/v1/query/%s/", ctx.DtableUuid)
	payload := strings.NewReader(fmt.Sprintf("{\"convert_keys\":%t,\"sql\":\"%s\"}", convertKeys, sql))
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[RowsWithSQL[map[string]interface{}]]{Request: req}
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
	data := make([]T, 0)
	util.ParseFromMapToAny(from.Results, data)
	to.Results = data
}

// GetRowByRowId 获取行数据
// convertKeys 确定列是作为其键 （true） 返回，还是作为其名称（默认为 false）返回。
func (api SeaTableApi) GetRowByRowId(ctx BaseContext, tableName string, rowId string, convertKeys bool) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/rows/%s/?table_name=%s&convert=%b", ctx.DtableUuid, rowId, tableName, convertKeys)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// InsertRowPosition 插入位置
type InsertRowPosition string

// InsertRowPosition 定义插入位置
const (
	// InsertBelow 在锚点行的下面插入
	InsertBelow InsertRowPosition = "insert_below"
	// InsertAbove 在锚点行的上面插入
	InsertAbove InsertRowPosition = "insert_above"
)

// InsertRowAnchor 插入锚点
type InsertRowAnchor struct {
	AnchorRowId string            // 瞄点行
	Position    InsertRowPosition // 插入行位置
}

// AddRow 添加行
func (api SeaTableApi) AddRow(ctx BaseContext, tableName string, data interface{}, anchor *InsertRowAnchor) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row": data}
	if anchor != nil && anchor.AnchorRowId != "" {
		body["anchor_row_id"] = anchor.AnchorRowId
		if anchor.Position != "" {
			body["row_insert_position"] = anchor.Position
		} else {
			body["row_insert_position"] = InsertBelow
		}
	}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// UpdateRow 更新行
func (api SeaTableApi) UpdateRow(ctx BaseContext, tableName string, rowId string, data interface{}) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row_id": rowId, "row": data}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// DeleteRow 删除行
func (api SeaTableApi) DeleteRow(ctx BaseContext, tableName string, rowId string) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row_id": rowId}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// AddRows 添加多行数据
func (api SeaTableApi) AddRows(ctx BaseContext, tableName string, data ...interface{}) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/batch-append-rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "rows": data}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// UpdateRows 更新多行数据
func (api SeaTableApi) UpdateRows(ctx BaseContext, tableName string, data ...struct {
	rowId string
	data  interface{}
}) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/batch-update-rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "updates": data}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// DeleteRows 删除多行数据
func (api SeaTableApi) DeleteRows(ctx BaseContext, tableName string, rowIds ...string) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/batch-delete-rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row_ids": rowIds}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// LockRows 锁定多行数据
func (api SeaTableApi) LockRows(ctx BaseContext, tableName string, rowIds ...string) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/lock-rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row_ids": rowIds}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}

// UnLockRows 解锁多行数据
func (api SeaTableApi) UnLockRows(ctx BaseContext, tableName string, rowIds ...string) SeaTableAction[map[string]interface{}] {
	url := api.wholeUrl("/dtable-server/api/v1/dtables/%s/unlock-rows/", ctx.DtableUuid)
	body := map[string]interface{}{"table_name": tableName, "row_ids": rowIds}
	payload := strings.NewReader(util.ParseToJsonString(body))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}
