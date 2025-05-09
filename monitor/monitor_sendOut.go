package monitor

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

// SendOutParam 定义订单参数的结构体
type SendOutParam struct {
	SysNum          string  `json:"sysNum"`                    // 系统内订单号
	PltNum          string  `json:"pltNum,omitempty"`          // 平台订单号
	OrderFee        string  `json:"orderFee,omitempty"`        // 订单金额
	PayTime         string  `json:"payTime"`                   // 支付时间
	DeliveryTime    string  `json:"deliveryTime,omitempty"`    // 发货时间
	DeliveryType    int     `json:"deliveryType,omitempty"`    // 发货模式
	ProDeliveryTime string  `json:"proDeliveryTime,omitempty"` // 承诺发货时间
	SourcePlatform  string  `json:"sourcePlatform"`            // 订单来源平台
	SourceReseller  string  `json:"sourceReseller,omitempty"`  // 订单来源分销商名称/自营
	StoreName       string  `json:"storeName,omitempty"`       // 系统内店铺名称
	WarehouseName   string  `json:"warehouseName,omitempty"`   // 发货仓/供应商名称
	ExpressNum      string  `json:"expressNum,omitempty"`      // 快递公司单号
	ExpressCode     string  `json:"expressCode,omitempty"`     // 快递公司的编码
	RemoteAreaFlag  string  `json:"remoteAreaFlag,omitempty"`  // 是否偏远地区
	Receiver        ManInfo `json:"receiver,omitempty"`        // 收件人信息
	Sender          ManInfo `json:"sender,omitempty"`          // 寄件人信息
}

func SendOut() {
	method := "sendOut"

	param := SendOutParam{
		SysNum:       "11111",
		ExpressNum:   "1111",
		ExpressCode:  "shunfeng",
		DeliveryTime: "2023-09-04 14:08:00",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := utils.DoMethodRequest(method, t, paramStr, config.MONITOR_ORDER_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
