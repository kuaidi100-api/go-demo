package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type BackOrderParam struct {
	Kuaidicom  string `json:"kuaidicom,omitempty"`
	Kuaidinum  string `json:"kuaidinum,omitempty"`
	PartnerId  string `json:"partnerId,omitempty"`
	PartnerKey string `json:"partnerKey,omitempty"`
	Phone      string `json:"phone,omitempty"`
	ImgType    int    `json:"imgType,omitempty"`
}

/*
*运单附件查询接口
 */
func BackOrder() {
	method := "backOrder"

	param := BackOrderParam{
		Kuaidicom:  "shunfeng",
		Kuaidinum:  "SF1234567",
		PartnerId:  "1234567",
		PartnerKey: "",
		Phone:      "13888888888",
		ImgType:    1,
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
