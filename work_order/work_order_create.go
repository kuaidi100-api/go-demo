package work_order

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type Attach struct {
	Type int    `json:"type"`
	Uri  string `json:"uri"`
}

type Desc struct {
	Attach []Attach `json:"attach"`
}

type WorkOrderCreateParam struct {
	Kuaidinum   string `json:"kuaidinum"`
	TelWeight   string `json:"telWeight"`
	CallBackUrl string `json:"callBackUrl"`
	SecondType  int    `json:"secondType"`
	Desc        Desc   `json:"desc"`
	Content     string `json:"content"`
}

/*
*创建工单
 */
func WorkOrderCreate() {

	attach := Attach{
		Type: 0,
		Uri:  "http://xxxxxxxxxxxxxxxxxxxxx",
	}

	desc := Desc{
		Attach: []Attach{attach},
	}

	param := WorkOrderCreateParam{
		Kuaidinum:   "asdsd123123123",
		TelWeight:   "1",
		CallBackUrl: "http://127.0.0.1:9100/apitest/apiOrder/callback",
		SecondType:  4,
		Desc:        desc,
		Content:     "重量异常",
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.WORK_ORDER_CREATE_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
