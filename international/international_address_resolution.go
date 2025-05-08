package international

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type InternationalAddressResolutionParam struct {
	Code     string `json:"code,omitempty"`
	Address  string `json:"address,omitempty"`
	Language string `json:"language,omitempty"`
}

func InternationalAddressResolution() {

	param := InternationalAddressResolutionParam{
		Code:     "US",
		Address:  "84 AlfordRd,GreatBarrington,MA01230,USA",
		Language: "en",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.INTERNATIONAL_ADDRESS_RESOLUTION_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
