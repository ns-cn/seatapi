package main

import (
	"fmt"
	"github.com/ns-cn/seatapi"
)

func main() {
	//ioApi := seatapi.SeaTableApi{Host: "https://cloud.seatable.io"}
	//action := ioApi.GetAccountToken("ns-cn@qq.com", "P@ssw0rd", "")
	//token, err := action.DoWithDefaultClient()
	//fmt.Println(token, err)
	ioApi := seatapi.SeaTableApi{Host: "https://cloud.seatable.io"}
	ioContextAction := ioApi.GetBaseContextWithApiToken("9b2ee0f10dc3563785f52554738f8cda52db557c")
	ioContext, _ := ioContextAction.DoWithDefaultClient()
	_rowsWithSQLAction := ioApi.ListRowsWithSQL(*ioContext, "select * from room", true)
	_result, _ := _rowsWithSQLAction.DoWithDefaultClient()
	fmt.Println(_result)
	infoAction := ioApi.GetBaseInfo(*ioContext)
	info, _ := infoAction.DoWithDefaultClient()
	fmt.Println(info)
	api := seatapi.SeaTableApi{Host: "https://cloud.seatable.cn"}
	baseContextWithApiTokenAction := api.GetBaseContextWithApiToken("4d59855f65ff7482fd071cece9b993c42de1b856")
	ctx, _ := baseContextWithApiTokenAction.DoWithDefaultClient()
	rowsWithSQLAction := api.ListRowsWithSQL(*ctx, "select * from room", true)
	result, err := rowsWithSQLAction.DoWithDefaultClient()
	fmt.Println(err)
	fmt.Println(result)
	s := seatapi.RowsWithSQL[map[string]interface{}]{}
	seatapi.ParseRowsWithSQL(*result, &s)
	fmt.Println(s)
	row := api.AddRow(*ctx, "t", map[string]interface{}{"名称": "测试3", "remark": "测试"}, nil)
	client, err := row.DoWithDefaultClient()
	fmt.Println(client)
	if err != nil {
		return
	}
}
