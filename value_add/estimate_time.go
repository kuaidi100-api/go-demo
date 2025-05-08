package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type EstimateTimeParam struct {
	Kuaidicom string `json:"kuaidicom,omitempty"`
	From      string `json:"from,omitempty"`
	To        string `json:"to,omitempty"`
	OrderTime string `json:"orderTime,omitempty"`
	ExpType   string `json:"expType,omitempty"`
}

/*
*快递预估时效查询接口
 */
func EstimateTime() {
	method := "time"

	param := EstimateTimeParam{
		Kuaidicom: "shunfeng",
		From:      "广东深圳南山区",
		To:        "北京海淀区",
		OrderTime: "2023-08-08 08:08:08",
		ExpType:   "标准快递",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.LABEL_ORDER_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
