package elec

import (
	"encoding/json"
	"fmt"
	"time"
)

func AuthThird() {

	param := Auth{
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
	_, err := DoAuthRequest(t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
