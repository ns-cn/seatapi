package seatapi

type BaseContext struct {
	AccessToken  string `json:"access_token"`  // access token
	WorkspaceID  int    `json:"workspace_id"`  // 工作空间ID
	AppName      string `json:"app_name"`      // 应用名
	DtableServer string `json:"dtable_server"` // dtable server
	DtableName   string `json:"dtable_name"`   // dtable name
	DtableSocket string `json:"dtable_socket"` // dtable socket
	DtableUuid   string `json:"dtable_uuid"`   // dtable uuid
	DtableDb     string `json:"dtable_db"`     // dtable db
}

/*
{
  "app_name": "api测试",
  "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJleHAiOjE3MDU2NDUzMDUsImR0YWJsZV91dWlkIjoiNjhmYTcwMjMtYWE1My00NjdkLWIwMDAtYjJlZGJlODQ0YTM0IiwidXNlcm5hbWUiOiIiLCJwZXJtaXNzaW9uIjoicnciLCJhcHBfbmFtZSI6ImFwaVx1NmQ0Ylx1OGJkNSJ9.O0ymWG_wsr0SX-o3U9monFK-zl14C4Or20CwjPWbpBc",
  "dtable_uuid": "68fa7023-aa53-467d-b000-b2edbe844a34",
  "dtable_server": "https://cloud.seatable.io/dtable-server/",
  "dtable_socket": "https://cloud.seatable.io/",
  "dtable_db": "https://cloud.seatable.io/dtable-db/",
  "workspace_id": 47529,
  "dtable_name": "test"
}
*/
