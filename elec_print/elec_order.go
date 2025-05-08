package elec_print

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
*电子面单下单接口
 */
func Order() {
	method := "order"

	recMan := Contact{
		Name:      "张三",
		Mobile:    "13888888888",
		PrintAddr: "广东深圳市南山区金蝶软件园",
		Company:   "",
	}

	sendMan := Contact{
		Name:      "李四",
		Mobile:    "13888888888",
		PrintAddr: "广东深圳市南山区金蝶软件园",
		Company:   "",
	}

	param := Param{
		PartnerId:           "123456",
		PartnerKey:          "",
		Code:                "",
		Kuaidicom:           "zhaijisong",
		RecMan:              recMan,
		SendMan:             sendMan,
		Cargo:               "test",
		TempId:              "60f6c17c7c223700131d8bc3",
		ChildTempId:         "61bff9efc66fb00013a1b168",
		BackTempId:          "61bffa26c66fb00013a1b16c",
		PayType:             "SHIPPER",
		ExpType:             "标准快递",
		Remark:              "测试下单,请勿发货",
		Collection:          "0",
		NeedChild:           "0",
		NeedBack:            "0",
		Count:               1,
		PrintType:           "CLOUD",
		Siid:                "KX100*******",
		NeedDesensitization: true,
		NeedLogo:            true,
		NeedOcr:             false,
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
