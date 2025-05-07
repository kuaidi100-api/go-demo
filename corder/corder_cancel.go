package corder

import (
	"encoding/json"
	"fmt"
	"time"
)

func Cancel() {
	method := "cancel"

	param := Param{
		TaskId:    "9FC293CA417E431F33046E64F4C4EC20",
		OrderId:   "20066771",
		CancelMsg: "内部测试单",
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
