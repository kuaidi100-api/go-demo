package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type OrderTaskParam struct {
	ShopType    string `json:"shopType,omitempty"`
	ShopId      string `json:"shopId,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
	UpdateAtMin string `json:"updateAtMin,omitempty"`
	UpdateAtMax string `json:"updateAtMax,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	Salt        string `json:"salt,omitempty"`
}

/*
*提交销售订单获取任务接口
 */
func OrderTask() {

	param := OrderTaskParam{
		ShopType:    "TAOBAO",
		ShopId:      "123",
		OrderStatus: "UNPAY",
		UpdateAtMin: "2025-05-06 11:11:11",
		UpdateAtMax: "2025-06-06 11:11:11",
		CallbackUrl: "www.baidu.com",
		Salt:        "abc",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.ORDER_TASK_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
