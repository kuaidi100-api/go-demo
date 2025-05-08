package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
)

type DetOCRParam struct {
	EnableTilt bool     `json:"enableTilt"`
	Include    []string `json:"include"`
	Image      *string  `json:"image"`
	ImageUrl   string   `json:"imageUrl"`
	PdfUrl     *string  `json:"pdfUrl"`
	HTMLUrl    *string  `json:"htmlUrl"`
}

/*
*快递面单OCR识别接口
 */
func DetOCR() {

	param := DetOCRParam{
		EnableTilt: false,
		Include:    []string{"barcode", "receiver", "sender"},
		Image:      nil,
		ImageUrl:   "https://cdn.kuaidi100.com/images/openapi/document/ocr_tem.png",
		PdfUrl:     nil,
		HTMLUrl:    nil,
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	m := map[string]string{
		"key":   config.KEY,
		"param": paramStr,
	}

	// 发送请求
	_, err := utils.DoMapRequest(m, config.DET_OCR_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
