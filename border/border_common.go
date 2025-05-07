package border

import (
	"go-test/config"
	"go-test/utils"
)

type Param struct {
	Kuaidicom        string `json:"kuaidicom,omitempty"`
	RecManName       string `json:"recManName,omitempty"`
	RecManMobile     string `json:"recManMobile,omitempty"`
	RecManPrintAddr  string `json:"recManPrintAddr,omitempty"`
	SendManName      string `json:"sendManName,omitempty"`
	SendManMobile    string `json:"sendManMobile,omitempty"`
	SendManPrintAddr string `json:"sendManPrintAddr,omitempty"`
	Cargo            string `json:"cargo,omitempty"`
	CallBackUrl      string `json:"callBackUrl,omitempty"`
	Payment          string `json:"payment,omitempty"`
	ServiceType      string `json:"serviceType,omitempty"`
	Weight           string `json:"weight,omitempty"`
	Remark           string `json:"remark,omitempty"`
	Salt             string `json:"salt,omitempty"`
	DayType          string `json:"dayType,omitempty"`
	PickupStartTime  string `json:"pickupStartTime,omitempty"`
	PickupEndTime    string `json:"pickupEndTime,omitempty"`
	PasswordSigning  string `json:"passwordSigning,omitempty"`
	ValinsPay        string `json:"valinsPay,omitempty"`
	Op               string `json:"op,omitempty"`
	PollCallBackUrl  string `json:"pollCallBackUrl,omitempty"`
	Resultv2         string `json:"resultv2,omitempty"`

	TaskId    string `json:"taskId,omitempty"`
	OrderId   string `json:"orderId,omitempty"`
	CancelMsg string `json:"cancelMsg,omitempty"`
}

func DoBorderRequest(method string, t string, param string) (string, error) {
	return utils.DoMethodRequest(method, t, param, config.B_ORDER_URL)
}
