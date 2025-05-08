package elec_print

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type ShopAuthorizeParam struct {
	ShopType    string `json:"shopType,omitempty"`
	CallbackUrl string `json:"callbackUrl,omitempty"`
	Salt        string `json:"salt,omitempty"`
}

/*
*获取店铺授权超链接接口
 */
func ShopAuthorize() {

	param := ShopAuthorizeParam{
		ShopType:    "TAOBAO",
		CallbackUrl: "www.baidu.com",
		Salt:        "taobao",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.SHOP_AUTHORIZE_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
