package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type EstimatePriceParam struct {
	SendAddr  string `json:"sendAddr"`
	RecAddr   string `json:"recAddr"`
	Kuaidicom string `json:"kuaidicom"`
	Weight    string `json:"weight"`
}

/*
*快递预估价格查询接口
 */
func EstimatePrice() {
	method := "price"

	param := EstimatePriceParam{
		SendAddr:  "深圳南山区",
		RecAddr:   "北京海淀区",
		Kuaidicom: "jd",
		Weight:    "12",
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
