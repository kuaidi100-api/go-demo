package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type PrintTaskParam struct {
	Siid string `json:"siid,omitempty"`
}

/*
*硬件状态查询接口
 */
func PrintTask() {
	method := "devstatus"

	param := PrintTaskParam{
		Siid: "12345678",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.PRINT_TASK_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
