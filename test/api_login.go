package main

import (
	"fmt"
	"seatapi"
)

func main() {
	api := seatapi.SeaTableApi{Host: "https://cloud.seatable.io"}
	action := api.GetAccountToken("ns-cn@qq.com", "P@ssw0rd", "")
	token, err := action.DoWithDefaultClient()
	fmt.Println(token, err)
	getBaseContextWithAccountTokenAction := api.GetBaseContextWithAccountToken(47529, "test", token.Token)
	baseContext, err := getBaseContextWithAccountTokenAction.DoWithDefaultClient()
	baseContext.WorkspaceID = 47529
	baseContext.DtableName = "test"
	if err != nil {
		return
	}
	fmt.Println(baseContext)
}
