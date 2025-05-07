package corder

import (
	"encoding/json"
	"fmt"
	"time"
)

func Create() {
	method := "cOrder"

	param := Param{
		Kuaidicom:        "yuantong",
		RecManName:       "王超",
		RecManMobile:     "12345678910",
		RecManPrintAddr:  "西藏日喀则市定日县珠穆朗玛峰",
		SendManName:      "王大",
		SendManMobile:    "12345678910",
		SendManPrintAddr: "西藏日喀则市定日县珠穆朗玛峰",
		Cargo:            "文件",
		CallBackUrl:      "http://www.baidu.com",
		Payment:          "SHIPPER",
		Weight:           "1",
		Remark:           "",
		Salt:             "",
		DayType:          "",
		PickupStartTime:  "",
		PickupEndTime:    "",
		Op:               "0",
		PollCallBackUrl:  "",
		Resultv2:         "0",
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
