package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type AuthThirdParam struct {
	Net         string `json:"net,omitempty"`
	CallBackUrl string `json:"callBackUrl,omitempty"`
	PartnerId   string `json:"partnerId,omitempty"`
	View        string `json:"view,omitempty"`
}

/*
*第三方平台账号授权
 */
func AuthThird() {

	param := AuthThirdParam{
		Net:         "123",
		CallBackUrl: "www.baidu.com",
		PartnerId:   "123",
		View:        "web",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.AUTH_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
