package border

import (
	"encoding/json"
	"fmt"
	"time"
)

func Price() {
	method := "price"

	param := Param{
		Kuaidicom:        "yuantong",
		RecManPrintAddr:  "广东深圳市深圳市南山区科技南十二路2号金蝶软件园",
		SendManPrintAddr: "广东省广州市华景软件园",
		ServiceType:      "标准快递",
		Weight:           "1",
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoBorderRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
