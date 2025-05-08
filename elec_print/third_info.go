package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type QueryBalanceParam struct {
	PartnerId  string `json:"partnerId,omitempty"`
	PartnerKey string `json:"partnerKey,omitempty"`
	Net        string `json:"net,omitempty"`
	Com        string `json:"com,omitempty"`
}

/*
*
第三方平台网点&面单余额接口
*/
func GetThirdInfo() {

	param := QueryBalanceParam{
		PartnerId:  "123",
		PartnerKey: "123",
		Net:        "taobao",
		Com:        "shentong",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.THIRD_INFO_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
