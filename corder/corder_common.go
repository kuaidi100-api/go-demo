package corder

import (
	"go-test/utils"
)

type Param struct {
	Kuaidicom        string `json:"kuaidicom"`
	RecManName       string `json:"recManName"`
	RecManMobile     string `json:"recManMobile"`
	RecManPrintAddr  string `json:"recManPrintAddr"`
	SendManName      string `json:"sendManName"`
	SendManMobile    string `json:"sendManMobile"`
	SendManPrintAddr string `json:"sendManPrintAddr"`
	Cargo            string `json:"cargo"`
	CallBackUrl      string `json:"callBackUrl"`
	Payment          string `json:"payment"`
	Weight           string `json:"weight"`
	Remark           string `json:"remark"`
	Salt             string `json:"salt"`
	DayType          string `json:"dayType"`
	PickupStartTime  string `json:"pickupStartTime"`
	PickupEndTime    string `json:"pickupEndTime"`
	Op               string `json:"op"`
	PollCallBackUrl  string `json:"pollCallBackUrl"`
	Resultv2         string `json:"resultv2"`

	TaskId    string `json:"taskId"`
	OrderId   string `json:"orderId"`
	CancelMsg string `json:"cancelMsg"`
}

const (
	C_ORDER_URL = "https://order.kuaidi100.com/order/corderapi.do" // 请求地址
)

func DoCorderRequest(method string, t string, param string) (string, error) {
	return utils.DoRequest(method, t, param, C_ORDER_URL)
}
