package work_order

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type WorkOrderQueryParam struct {
	ConsultId string `json:"consultId,omitempty"`
}

/*
*查询工单
 */
func WorkOrderQuery() {

	param := WorkOrderQueryParam{
		ConsultId: "1",
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.WORK_ORDER_QUERY_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
