package seatapi

import (
	"net/http"
)

// GetBaseInfo 获取基础信息
func (api SeaTableApi) GetBaseInfo(ctx *BaseContext) SeaTableAction[BaseInfo] {
	url := api.assignUrl(ctx.DtableServer, "/dtables/%s", ctx.DtableUuid)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", api.tokenHeader(ctx.AccessToken))
	return SeaTableAction[BaseInfo]{Request: req}
}
