package base

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
)

type PollMapParameters struct {
	Callbackurl string `json:"callbackurl"`
	Salt        string `json:"salt"`
	Phone       string `json:"phone"`
	Resultv2    string `json:"resultv2"`
}

type PollMapParam struct {
	Company    string            `json:"company"`
	Number     string            `json:"number"`
	From       string            `json:"from"`
	To         string            `json:"to"`
	Key        string            `json:"key"`
	Parameters PollMapParameters `json:"parameters"`
}

/*
*地图轨迹推送接口
 */
func PollMap() {

	parameters := PollMapParameters{
		Callbackurl: "http://www.您的域名.com/kuaidi?callbackid=...",
		Salt:        "*",
		Phone:       "",
		Resultv2:    "5",
	}

	param := PollMapParam{
		Company:    "ems",
		Number:     "1136281381675",
		From:       "广东省深圳市南山区",
		To:         "北京市朝阳区",
		Key:        config.KEY,
		Parameters: parameters,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	m := map[string]string{
		"schema": "json",
		"param":  paramStr,
	}

	// 发送请求
	_, err := utils.DoMapRequest(m, config.POLL_MAP_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
