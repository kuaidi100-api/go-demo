package elec_print

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
*自定义模板打印接口
 */
func custom() {
	method := "custom"

	customParam := CustomParam{
		Name:       "12213",
		Flag:       "出库",
		Count:      111,
		Unit:       "货柜",
		Total:      "2000",
		OrderId:    "8888888888888",
		CustomerId: "66666666",
		Cargo1:     "苹果",
		Cargo2:     "玉米",
		Num:        "SF1536245218562",
		Label:      "拼多多",
		QrCode:     "6666666666666",
		Org:        "快递100",
	}

	// 初始化参数结构体
	param := Param{
		TempId:      "476f6f769e57447fb84398eefae2ae28",
		PrintType:   "CLOUD",
		Siid:        "KX100siid",
		CustomParam: customParam,
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
