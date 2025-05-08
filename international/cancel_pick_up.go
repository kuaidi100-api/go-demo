package international

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type CancelPickUpParam struct {
	PickUpDate               string `json:"pickUpDate"`
	Location                 string `json:"location"`
	PartnerId                string `json:"partnerId"`
	PartnerName              string `json:"partnerName"`
	PartnerSecret            string `json:"partnerSecret"`
	PartnerKey               string `json:"partnerKey"`
	Code                     string `json:"code"`
	Kuaidicom                string `json:"kuaidicom"`
	PickupConfirmationNumber string `json:"pickupConfirmationNumber"`
}

/*
*国际电子面单取消预约API
 */
func CancelPickUp() {

	param := CancelPickUpParam{
		PickUpDate:               "2022-05-31",
		Location:                 "CN",
		PartnerId:                "xxx",
		PartnerName:              "xxx",
		PartnerSecret:            "xxx",
		PartnerKey:               "xxx",
		Code:                     "xxx",
		Kuaidicom:                "dhl",
		PickupConfirmationNumber: "CBJ220608002901",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.CANCEL_PICK_UP_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
