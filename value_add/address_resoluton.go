package value_add

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type AddressResolutionParam struct {
	Content  string `json:"content,omitempty"`
	Image    string `json:"image,omitempty"`
	ImageUrl string `json:"imageUrl,omitempty"`
	PdfUrl   string `json:"pdfUrl,omitempty"`
	HtmlUrl  string `json:"htmlUrl,omitempty"`
}

/*
*智能地址解析接口
 */
func AddressResolution() {

	param := AddressResolutionParam{
		Content:  "张三广东省深圳市南山区粤海街道科技南十二路金蝶软件园13088888888",
		Image:    "qweasftefds",
		ImageUrl: "www.imageurl.png",
		PdfUrl:   "www.pdfurl.pdf",
		HtmlUrl:  "www.htmlurl.html",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.ADDRESS_RESOLUTION_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
