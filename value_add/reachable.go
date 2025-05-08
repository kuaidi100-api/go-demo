package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type ReachableParam struct {
	RecManPrintAddr  string `json:"recManPrintAddr"`
	SendManPrintAddr string `json:"sendManPrintAddr"`
	RecManMobile     string `json:"recManMobile"`
	SendManName      string `json:"sendManName"`
	RecManName       string `json:"recManName"`
	Kuaidicom        string `json:"kuaidicom"`
	SendManMobile    string `json:"sendManMobile"`
}

/*
*快递可用性接口
 */
func Reachable() {
	method := "reachable"

	param := ReachableParam{
		RecManPrintAddr:  "浙江省湖州市吴兴区织****",
		SendManPrintAddr: "福建省宁德市霞***",
		RecManMobile:     "****",
		SendManName:      "****",
		RecManName:       "***",
		Kuaidicom:        "yuantong",
		SendManMobile:    "*****",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.REACHABLE_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
