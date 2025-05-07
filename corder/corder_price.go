package corder

import (
	"encoding/json"
	"fmt"
	"time"
)

func Price() {
	method := "price"

	param := Param{
		Kuaidicom:        "shunfeng",
		SendManPrintAddr: "北京市海淀区颐和园路5号",
		RecManPrintAddr:  "北京市海淀区双清路30号",
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoCorderRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
