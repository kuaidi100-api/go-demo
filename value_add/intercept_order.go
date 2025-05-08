package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type RecManInfo struct {
	Name      string `json:"name,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	PrintAddr string `json:"printAddr,omitempty"`
}

type InterceptOrderParam struct {
	OrderId       string     `json:"orderId,omitempty"`
	Kuaidicom     string     `json:"kuaidicom,omitempty"`
	Kuaidinum     string     `json:"kuaidinum,omitempty"`
	PartnerId     string     `json:"partnerId,omitempty"`
	PartnerKey    string     `json:"partnerKey,omitempty"`
	PartnerSecret string     `json:"partnerSecret,omitempty"`
	InterceptType string     `json:"interceptType,omitempty"`
	Code          string     `json:"code,omitempty"`
	Net           string     `json:"net,omitempty"`
	Reason        string     `json:"reason,omitempty"`
	RecManInfo    RecManInfo `json:"recManInfo,omitempty"`
}

/*
*拦截改址接口
 */
func InterceptOrder() {
	method := "interceptOrder"

	recManInfo := RecManInfo{
		Name:      "测试",
		Mobile:    "138888888888",
		PrintAddr: "广东深圳市南山区金蝶软件园",
	}

	param := InterceptOrderParam{
		OrderId:       "",
		Kuaidicom:     "debangkuaidi",
		Kuaidinum:     "2222",
		PartnerId:     "22222",
		PartnerKey:    "",
		PartnerSecret: "",
		InterceptType: "MODIFY_ADDR",
		Code:          "",
		Net:           "",
		Reason:        "",
		RecManInfo:    recManInfo,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.LABEL_ORDER_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
