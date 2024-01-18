package main

import (
	"fmt"
	"github.com/ns-cn/seatapi"
)

func main() {
	api := seatapi.SeaTableApi{Host: "https://cloud.seatable.io"}
	action := api.GetAccountToken("ns-cn@qq.com", "P@ssw0rd", "")
	token, err := action.DoWithDefaultClient()
	fmt.Println(token, err)
	baseContextWithApiTokenAction := api.GetBaseContextWithApiToken("9b2ee0f10dc3563785f52554738f8cda52db557c")
	ctx, _ := baseContextWithApiTokenAction.DoWithDefaultClient()
	rowsWithSQLAction := api.ListRowsWithSQL(*ctx, "select * from room", false)
	result, _ := rowsWithSQLAction.DoWithDefaultClient()
	fmt.Println(result)
	row := api.AddRow(*ctx, "t", map[string]interface{}{"名称": "测试3", "remark": "测试"}, nil)
	client, err := row.DoWithDefaultClient()
	fmt.Println(client)
	if err != nil {
		return
	}
}
