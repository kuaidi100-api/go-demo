package international

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type PickUpParam struct {
	PickupTimestamp         string        `json:"pickupTimestamp"`
	PickupLocationCloseTime string        `json:"pickupLocationCloseTime"`
	Instruction             string        `json:"instruction"`
	PartnerId               string        `json:"partnerId"`
	PartnerName             string        `json:"partnerName"`
	PartnerSecret           string        `json:"partnerSecret"`
	Kuaidicom               string        `json:"kuaidicom"`
	SendMan                 SendMan       `json:"sendMan"`
	RecMan                  RecMan        `json:"recMan"`
	PackageInfos            []PackageInfo `json:"packageInfos"`
}

/*
*国际电子面单下单API
 */
func PickUp() {
	sendMan := SendMan{
		Name:        "Kaka",
		Mobile:      "13500000000",
		Addr:        "Kingdee Software Park",
		CountryCode: "CN",
		City:        "SHEN ZHEN",
		Zipcode:     "518057",
	}

	recMan := RecMan{
		Name:        "Cindy Martinez / Ana Luz Medina",
		Mobile:      "(86)13510002000",
		Addr:        "Apoquindo 4001, of. 501. Las Condes",
		CountryCode: "CL",
		City:        "Santiago",
		Zipcode:     "7550000",
	}

	packageInfos := []PackageInfo{
		{
			Height: "10",
			Width:  "20",
			Length: "11",
			Weight: "0.1",
		},
		{
			Height: "10",
			Width:  "20",
			Length: "11",
			Weight: "0.1",
		},
		{
			Height: "10",
			Width:  "20",
			Length: "11",
			Weight: "0.1",
		},
	}

	param := PickUpParam{
		PickupTimestamp:         "2022-05-16 08:41:35",
		PickupLocationCloseTime: "17:00",
		Instruction:             "Collect at reception",
		PartnerId:               "******",
		PartnerName:             "******",
		PartnerSecret:           "******",
		Kuaidicom:               "dhl",
		SendMan:                 sendMan,
		RecMan:                  recMan,
		PackageInfos:            packageInfos,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.PICK_UP_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
