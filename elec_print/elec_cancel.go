package elec_print

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
*
电子面单取消接口
*/
func Cancel() {
	method := "cancel"

	param := Param{
		PartnerId:     "123",
		PartnerKey:    "123",
		PartnerSecret: "123",
		PartnerName:   "123",
		Net:           "123",
		Code:          "123",
		Kuaidicom:     "******",
		Kuaidinum:     "******",
		OrderId:       "123",
		CheckMan:      "123",
		Reason:        "123",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoLabelOrderRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
