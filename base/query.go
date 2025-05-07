package base

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
)

type QueryParam struct {
	Com      string `json:"com"`
	Num      string `json:"num"`
	Phone    string `json:"phone"`
	From     string `json:"from"`
	To       string `json:"to"`
	Resultv2 string `json:"resultv2"`
	Show     string `json:"show"`
	Order    string `json:"order"`
	LANG     string `json:"lang"`
}

/*
实时快递查询接口
*/
func Query() {

	param := QueryParam{
		Com:      "yuantong",
		Num:      "em263999513jp",
		Phone:    "13868688888",
		From:     "广东省深圳市南山区",
		To:       "北京市朝阳区",
		Resultv2: "4",
		Show:     "0",
		Order:    "desc",
		LANG:     "zh",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 发送请求
	_, err := utils.CustomerRequest(paramStr, config.QUERY_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
