package elec

import (
	"encoding/json"
	"fmt"
	"time"
)

func QueryBalance() {

	param := Param{
		PartnerId:  "123",
		PartnerKey: "123",
		Net:        "taobao",
		Com:        "shentong",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoQueryBalanceRequest(t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
