package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type LogisticSendParam struct {
	ShopType     string `json:"shopType,omitempty"`
	ShopId       string `json:"shopId,omitempty"`
	OrderNum     string `json:"orderNum,omitempty"`
	KuaidiCom    string `json:"KuaidiCom,omitempty"`
	KuaidiNum    string `json:"KuaidiNum,omitempty"`
	SubOrderNums string `json:"SubOrderNums,omitempty"`
}

/*
*提交售后（退货）订单获取任务接口
 */
func LogisticSend() {

	param := LogisticSendParam{
		ShopType:     "TAOBAO",
		ShopId:       "123",
		OrderNum:     "12345",
		KuaidiCom:    "shunfeng",
		KuaidiNum:    "SF54321",
		SubOrderNums: "",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.LOGISTIC_SEND_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
