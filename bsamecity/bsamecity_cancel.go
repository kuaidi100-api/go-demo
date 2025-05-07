package bsamecity

import (
	"encoding/json"
	"fmt"
	"time"
)

func Cancel() {
	method := "cancel"

	// 构建请求参数
	param := Param{
		OrderId:       "100241",
		CancelMsgType: 1,
		CancelMsg:     "测试寄件",
		TaskId:        "DE49A5C45B0845328CE0AE8EF9951EC5",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoBsamecityRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
