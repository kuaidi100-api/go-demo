package bsamecity

import (
	"go-test/utils"
)

type Param struct {
	Kuaidicom        string  `json:"kuaidicom,omitempty"`
	LbsType          int     `json:"lbsType,omitempty"`
	RecManName       string  `json:"recManName,omitempty"`
	RecManMobile     string  `json:"recManMobile,omitempty"`
	RecManProvince   string  `json:"recManProvince,omitempty"`
	RecManCity       string  `json:"recManCity,omitempty"`
	RecManDistrict   string  `json:"recManDistrict,omitempty"`
	RecManAddr       string  `json:"recManAddr,omitempty"`
	RecManLat        string  `json:"recManLat,omitempty"`
	RecManLng        string  `json:"recManLng,omitempty"`
	SendManName      string  `json:"sendManName,omitempty"`
	SendManMobile    string  `json:"sendManMobile,omitempty"`
	SendManProvince  string  `json:"sendManProvince,omitempty"`
	SendManCity      string  `json:"sendManCity,omitempty"`
	SendManDistrict  string  `json:"sendManDistrict,omitempty"`
	SendManAddr      string  `json:"sendManAddr,omitempty"`
	SendManLat       string  `json:"sendManLat,omitempty"`
	SendManLng       string  `json:"sendManLng,omitempty"`
	Weight           string  `json:"weight,omitempty"`
	Remark           string  `json:"remark,omitempty"`
	Volume           string  `json:"volume,omitempty"`
	Salt             string  `json:"salt,omitempty"`
	CallbackUrl      string  `json:"callbackUrl,omitempty"`
	OrderType        int     `json:"orderType,omitempty"`
	ExpectPickupTime string  `json:"expectPickupTime,omitempty"`
	ExpectFinishTime string  `json:"expectTimeFinish,omitempty"`
	Insurance        string  `json:"insurance,omitempty"`
	Price            string  `json:"price,omitempty"`
	Goods            []Goods `json:"goods,omitempty"`

	OrderId       string `json:"orderId,omitempty"`
	CancelMsgType int    `json:"cancelMsgType,omitempty"`
	CancelMsg     string `json:"cancelMsg,omitempty"`
	TaskId        string `json:"taskId,omitempty"`

	Tips string `json:"tips,omitempty"`
}

type Goods struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Count int    `json:"count,omitempty"`
}

const (
	BSAMECITY_URL = "https://api.kuaidi100.com/bsamecity/order" // 请求地址
)

func DoBsamecityRequest(method string, t string, param string) (string, error) {
	return utils.DoRequest(method, t, param, BSAMECITY_URL)
}
