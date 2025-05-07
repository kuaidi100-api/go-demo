package elec

import (
	"go-test/utils"
)

type Contact struct {
	Name      string `json:"name,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	PrintAddr string `json:"printAddr,omitempty"`
	Company   string `json:"company,omitempty"`
}

type CustomParam struct {
	Name       string `json:"name"`
	Flag       string `json:"flag"`
	Count      int    `json:"count"`
	Unit       string `json:"unit"`
	Total      string `json:"total"`
	OrderId    string `json:"orderId"`
	CustomerId string `json:"customerId"`
	Cargo1     string `json:"cargo1"`
	Cargo2     string `json:"cargo2"`
	Num        string `json:"num"`
	Label      string `json:"label"`
	QrCode     string `json:"qrCode"`
	Org        string `json:"org"`
}

// 定义参数结构体
type Param struct {
	PartnerId           string  `json:"partnerId,omitempty"`
	PartnerKey          string  `json:"partnerKey,omitempty"`
	Code                string  `json:"code,omitempty"`
	Kuaidicom           string  `json:"kuaidicom,omitempty"`
	RecMan              Contact `json:"recMan,omitempty"`
	SendMan             Contact `json:"sendMan,omitempty"`
	Cargo               string  `json:"cargo,omitempty"`
	TempId              string  `json:"tempId,omitempty"`
	ChildTempId         string  `json:"childTempId,omitempty"`
	BackTempId          string  `json:"backTempId,omitempty"`
	PayType             string  `json:"payType,omitempty"`
	ExpType             string  `json:"expType,omitempty"`
	Remark              string  `json:"remark,omitempty"`
	Collection          string  `json:"collection,omitempty"`
	NeedChild           string  `json:"needChild,omitempty"`
	NeedBack            string  `json:"needBack,omitempty"`
	Count               int     `json:"count,omitempty"`
	PrintType           string  `json:"printType,omitempty"`
	Siid                string  `json:"siid,omitempty"`
	NeedDesensitization bool    `json:"needDesensitization,omitempty"`
	NeedLogo            bool    `json:"needLogo,omitempty"`
	NeedOcr             bool    `json:"needOcr,omitempty"`

	TaskId        string `json:"taskId,omitempty"`
	PartnerSecret string `json:"partnerSecret,omitempty"`
	PartnerName   string `json:"partnerName,omitempty"`
	Net           string `json:"net,omitempty"`
	Kuaidinum     string `json:"kuaidinum,omitempty"`
	OrderId       string `json:"orderId,omitempty"`
	CheckMan      string `json:"checkMan,omitempty"`
	Reason        string `json:"reason,omitempty"`

	CustomParam CustomParam `json:"customParam"`

	Com string `json:"com,omitempty"`
}

type Auth struct {
	Net         string `json:"net,omitempty"`
	CallBackUrl string `json:"callBackUrl,omitempty"`
	PartnerId   string `json:"partnerId,omitempty"`
	View        string `json:"view,omitempty"`
}

const (
	AUTH_URL          = "https://poll.kuaidi100.com/printapi/authThird.do" // 第三方平台账号授权
	LABEL_ORDER_URL   = "https://api.kuaidi100.com/label/order"            // 自定义模板打印、自定义模板打印复打、电子面单下单/复打/取消请求地址
	QUERY_BALANCE_URL = "http://poll.kuaidi100.com/eorderapi.do"           // 第三方平台网点&面单余额接口
	ELEC_OCR_URL      = "http://api.kuaidi100.com/elec/detocr"             // 电子面单OCR识别请求地址
)

func DoAuthRequest(t string, param string) (string, error) {
	return utils.DoRequest("", t, param, AUTH_URL)
}

func DoLabelOrderRequest(method string, t string, param string) (string, error) {
	return utils.DoRequest(method, t, param, LABEL_ORDER_URL)
}

func DoQueryBalanceRequest(t string, param string) (string, error) {
	return utils.DoRequest("", t, param, QUERY_BALANCE_URL)
}

// DoElecOcrRequest 处理电子面单OCR识别请求
func DoElecOcrRequest(param string) (string, error) {
	return utils.DoRequest("", "", param, ELEC_OCR_URL)
}
