package international

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type SendMan struct {
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	Addr        string `json:"addr"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Company     string `json:"company"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	VatNum      string `json:"vatNum"`
	EoriNum     string `json:"eoriNum"`
	IossNum     string `json:"iossNum"`
}

type RecMan struct {
	Name        string `json:"name"`
	Mobile      string `json:"mobile"`
	Addr        string `json:"addr"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Company     string `json:"company"`
	CountryCode string `json:"countryCode"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	Tel         string `json:"tel"`
	Email       string `json:"email"`
	VatNum      string `json:"vatNum"`
	EoriNum     string `json:"eoriNum"`
	IossNum     string `json:"iossNum"`
}

type PackageInfo struct {
	Height           string `json:"height"`
	Width            string `json:"width"`
	Length           string `json:"length"`
	Weight           string `json:"weight"`
	PackageReference string `json:"packageReference"`
}

type ExportInfo struct {
	ZhName                    string  `json:"zhName"`
	EnName                    string  `json:"enName"`
	NetWeight                 float64 `json:"netWeight"`
	GrossWeight               float64 `json:"grossWeight"`
	ManufacturingCountryCode  string  `json:"manufacturingCountryCode"`
	UnitPrice                 string  `json:"unitPrice"`
	Quantity                  float64 `json:"quantity"`
	QuantityUnitOfMeasurement string  `json:"quantityUnitOfMeasurement"`
	Desc                      string  `json:"desc"`
	ImportCommodityCode       string  `json:"importCommodityCode"`
	ExportCommodityCode       string  `json:"exportCommodityCode"`
	NumberOfPieces            int     `json:"numberOfPieces"`
}

type InvoiceInfo struct {
	Date      string  `json:"date"`
	Number    string  `json:"number"`
	Type      *string `json:"type"`
	Title     string  `json:"title"`
	Signature string  `json:"signature"`
	PltEnable bool    `json:"pltEnable"`
}

type DutiesPaymentType struct {
	PaymentType string `json:"paymentType"`
	Account     string `json:"account"`
}

type FreightPaymentType struct {
	PaymentType string `json:"paymentType"`
	Account     string `json:"account"`
}

type CustomsClearance struct {
	Purpose  string `json:"purpose"`
	Document bool   `json:"document"`
}

type ApiCallParam struct {
	PartnerId          string             `json:"partnerId"`
	PartnerName        string             `json:"partnerName"`
	PartnerSecret      string             `json:"partnerSecret"`
	Kuaidicom          string             `json:"kuaidicom"`
	SendMan            SendMan            `json:"sendMan"`
	RecMan             RecMan             `json:"recMan"`
	Cargo              string             `json:"cargo"`
	Count              string             `json:"count"`
	Weight             string             `json:"weight"`
	ExpType            string             `json:"expType"`
	Remark             string             `json:"remark"`
	PackageInfos       []PackageInfo      `json:"packageInfos"`
	ExportInfos        []ExportInfo       `json:"exportInfos"`
	CustomsValue       string             `json:"customsValue"`
	TradeTerm          string             `json:"tradeTerm"`
	NeedInvoice        bool               `json:"needInvoice"`
	InvoiceInfo        InvoiceInfo        `json:"invoiceInfo"`
	DutiesPaymentType  DutiesPaymentType  `json:"dutiesPaymentType"`
	FreightPaymentType FreightPaymentType `json:"freightPaymentType"`
	CustomsClearance   CustomsClearance   `json:"customsClearance"`
}

/*
*国际电子面单下单API
 */
func ApiCall() {

	sendMan := SendMan{
		Name:        "Kaka",
		Mobile:      "13500000000",
		Addr:        "Kingdee Software Park",
		District:    "Hi-tech Park,Nanshang District",
		Province:    "",
		Company:     "QIAN HAI BAI DI",
		CountryCode: "CN",
		City:        "SHEN ZHEN",
		Zipcode:     "518057",
		Tel:         "0755-5890123",
		Email:       "12344655@qq.com",
		VatNum:      "IOSS23249923",
		EoriNum:     "IOSS23249923",
		IossNum:     "IOSS23249923",
	}

	recMan := RecMan{
		Name:        "Cindy Martinez / Ana Luz Medina",
		Mobile:      "(86)13560312000",
		Addr:        "Apoquindo 4001, of. 501. Las Condes",
		District:    "Santiago, Chile",
		Province:    "",
		Company:     "Lamaignere Chile S.A.",
		CountryCode: "CL",
		City:        "Santiago",
		Zipcode:     "7550000",
		Tel:         " +56 (9) 1242355",
		Email:       "12344699@qq.com",
		VatNum:      "IOSS23249923",
		EoriNum:     "IOSS23249923",
		IossNum:     "IOSS23249923",
	}

	packageInfo := PackageInfo{
		Height:           "10",
		Width:            "20",
		Length:           "11",
		Weight:           "0.1",
		PackageReference: "just a user remark",
	}

	exportInfo := ExportInfo{
		ZhName:                    "",
		EnName:                    "",
		NetWeight:                 0.1,
		GrossWeight:               0.1,
		ManufacturingCountryCode:  "CN",
		UnitPrice:                 "10.0",
		Quantity:                  1.0,
		QuantityUnitOfMeasurement: "PCS",
		Desc:                      "test",
		ImportCommodityCode:       "6109100021",
		ExportCommodityCode:       "6109100021",
		NumberOfPieces:            1,
	}

	dutiesPaymentType := DutiesPaymentType{
		PaymentType: "DDU",
		Account:     "60*****43",
	}

	freightPaymentType := FreightPaymentType{
		PaymentType: "SHIPPER",
		Account:     "60*****43",
	}

	customsClearance := CustomsClearance{
		Purpose:  "",
		Document: true,
	}

	invoiceInfo := InvoiceInfo{
		Date:      "2021-08-12",
		Number:    "15462412",
		Type:      nil,
		Title:     "test",
		Signature: "base64Str",
		PltEnable: true,
	}

	param := ApiCallParam{
		PartnerId:          "12344",
		PartnerName:        "dsfgsgs",
		PartnerSecret:      "dfsfsfs",
		Kuaidicom:          "dhl",
		SendMan:            sendMan,
		RecMan:             recMan,
		Cargo:              "test don't ship",
		Count:              "1",
		Weight:             "0.1",
		ExpType:            "parcel-normal",
		Remark:             "just a test demo",
		PackageInfos:       []PackageInfo{packageInfo},
		ExportInfos:        []ExportInfo{exportInfo},
		CustomsValue:       "10.0",
		TradeTerm:          "DAP",
		NeedInvoice:        true,
		InvoiceInfo:        invoiceInfo,
		DutiesPaymentType:  dutiesPaymentType,
		FreightPaymentType: freightPaymentType,
		CustomsClearance:   customsClearance,
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoRequest(t, paramStr, config.API_CALL_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
