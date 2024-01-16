package seatapi

import (
	"fmt"
	"net/http"
	"strings"
)

// GetAccountToken 获取账号上下文
func (api SeaTableApi) GetAccountToken(username string, password string, xSeaFileOtp string) SeaTableAction[AccountToken] {
	url := api.wholeUrl("/api2/auth-token/")
	payload := strings.NewReader(fmt.Sprintf("username=%s&password=%s", username, password))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	if xSeaFileOtp != "" {
		req.Header.Add("X-SEAFILE-OTP", xSeaFileOtp)
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	action := SeaTableAction[AccountToken]{}
	action.Request = req
	return action
}

// AccountToken 获取账号token操作的响应类型
type AccountToken struct {
	Token string `json:"token"`
}

// GetBaseContextWithAccountToken 使用账号token获取Base上下文
// 该方式仅能获取到AccessToken,DtableUuid,DtableServer,DtableSocket,DtableDb
// 其他额外的字段请手动设置
func (api SeaTableApi) GetBaseContextWithAccountToken(workspaceId int, baseName string, accountToken string) SeaTableAction[BaseContext] {
	url := fmt.Sprintf("%s/api/v2.1/workspace/%d/dtable/%s/access-token/", api.Host, workspaceId, baseName)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", accountToken))
	action := SeaTableAction[BaseContext]{}
	action.Request = req
	return action
}

// GetBaseContextWithApiToken 使用ApiToken获取Base上下文
// 能获取到BaseContext全量数据
func (api SeaTableApi) GetBaseContextWithApiToken(apiToken string) SeaTableAction[BaseContext] {
	url := api.wholeUrl("/api/v2.1/dtable/app-access-token/")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", apiToken))
	action := SeaTableAction[BaseContext]{}
	action.Request = req
	return action
}
