package border

import (
	"encoding/json"
	"fmt"
	"time"
)

func Create() {
	method := "create"

	param := Param{
		Kuaidicom:        "yuantong",
		RecManName:       "王超",
		RecManMobile:     "13842569988",
		RecManPrintAddr:  "广东深圳市深圳市南山区科技南十二路2号金蝶软件园",
		SendManName:      "王大",
		SendManMobile:    "13842569988",
		SendManPrintAddr: "广东省广州市华景软件园",
		Cargo:            "文件",
		CallBackUrl:      "http://www.baidu.com",
		Payment:          "SHIPPER",
		ServiceType:      "快递标准",
		Weight:           "1",
		Remark:           "",
		Salt:             "",
		DayType:          "",
		PickupStartTime:  "",
		PickupEndTime:    "",
		PasswordSigning:  "Y",
		ValinsPay:        "",
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
	_, err := DoBorderRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
