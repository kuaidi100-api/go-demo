package work_order

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type WorkOrderReplyParam struct {
	ConsultId string   `json:"consultId,omitempty"`
	Content   string   `json:"content,omitempty"`
	Attach    []Attach `json:"attach"`
}

/*
*查询留言
 */
func WorkOrderReply() {
	method := "queryReply"

	param := WorkOrderReplyParam{
		ConsultId: "1056",
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.WORK_ORDER_REPLY_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}

/*
*
新增留言
*/
func WorkOrderAddReply() {
	method := "addReply"

	attach := Attach{
		Uri:  "xxxxx",
		Type: 0,
	}

	param := WorkOrderReplyParam{
		ConsultId: "1023",
		Content:   "testApi",
		Attach:    []Attach{attach},
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.WORK_ORDER_REPLY_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
