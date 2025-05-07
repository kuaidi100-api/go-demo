package bsamecity

import (
	"encoding/json"
	"fmt"
	"time"
)

func AddFee() {
	method := "addfee"

	// 构建请求参数
	param := Param{
		OrderId: "100213",
		Remark:  "",
		TaskId:  "0E1536182BAE416080AC3C5DE6896F03",
		Tips:    "10",
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
