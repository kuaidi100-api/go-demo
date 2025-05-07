package elec

import (
	"encoding/json"
	"fmt"
	"time"
)

func printOld() {
	method := "printOld"

	param := Param{
		Siid:   "KX100*******",
		TaskId: "1234",
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
