package seatapi

import (
	"fmt"
	"net/http"
	"strings"
)

// DuplicateTable 复制表
func (api SeaTableApi) DuplicateTable(ctx *BaseContext, tableName string, isDuplicateRecords bool) SeaTableAction[BaseTable] {
	url := api.assignUrl(ctx.DtableServer, "/api/v1/dtables/%s/tables/duplicate-table/", ctx.DtableUuid)
	payload := strings.NewReader(fmt.Sprintf("{\"is_duplicate_records\":%v,\"table_name\":\"%s\"}", isDuplicateRecords, tableName))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[BaseTable]{Request: req}
}

// RenameTable 重命名表
func (api SeaTableApi) RenameTable(ctx *BaseContext, tableName string, newTableName string) SeaTableAction[BaseTable] {
	url := api.assignUrl(ctx.DtableServer, "/api/v1/dtables/%s/tables/", ctx.DtableUuid)
	payload := strings.NewReader(fmt.Sprintf("{\"table_name\":\"%s\",\"new_table_name\":\"%s\"}", tableName, newTableName))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[BaseTable]{Request: req}
}

// DeleteTable 删除表
func (api SeaTableApi) DeleteTable(ctx *BaseContext, tableName string) SeaTableAction[map[string]interface{}] {
	url := api.assignUrl(ctx.DtableServer, "/api/v1/dtables/%s/tables/", ctx.DtableUuid)
	payload := strings.NewReader(fmt.Sprintf("{\"table_name\":\"%s\"}", tableName))
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[map[string]interface{}]{Request: req}
}
