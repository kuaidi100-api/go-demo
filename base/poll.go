package base

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
)

type PollParameters struct {
	Callbackurl string `json:"callbackurl"`
	Salt        string `json:"salt"`
	Phone       string `json:"phone"`
	Resultv2    string `json:"resultv2"`
}

type PollParam struct {
	Company    string         `json:"company"`
	Number     string         `json:"number"`
	Key        string         `json:"key"`
	Parameters PollParameters `json:"parameters"`
}

/*
快递信息推送服务-订阅接口
*/
func Poll() {

	parameters := PollParameters{
		Callbackurl: "http://www.您的域名.com/kuaidi?callbackid=...",
		Salt:        "XXXXXXXXXX",
		Phone:       "13868688888",
		Resultv2:    "4",
	}

	param := PollParam{
		Company:    "yuantong",
		Number:     "YT6186594166532",
		Key:        "XXX ",
		Parameters: parameters,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 发送请求
	_, err := utils.CustomerRequest(paramStr, config.POLL_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
