package monitor

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"time"
)

type ManInfo struct {
	Name    string `json:"name"`
	Mobile  string `json:"mobile"`
	Address string `json:"address"`
}

// OrderExportParam 定义订单参数的结构体
type OrderExportParam struct {
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

func OrderExport() {
	method := "orderExport"
	param := OrderExportParam{
		SysNum:          "SYS123456",
		PltNum:          "PLT789012",
		OrderFee:        "199.00",
		PayTime:         "2024-07-20 15:30:00",
		DeliveryTime:    "2024-07-21 09:00:00",
		DeliveryType:    0,
		ProDeliveryTime: "2024-07-22 18:00:00",
		SourcePlatform:  "JD",
		SourceReseller:  "自营",
		StoreName:       "官方旗舰店",
		WarehouseName:   "华北仓库",
		ExpressNum:      "JD1234567890",
		ExpressCode:     "jd",
		RemoteAreaFlag:  "0",
		Receiver: ManInfo{
			Name:    "张三",
			Mobile:  "13800138000",
			Address: "北京市朝阳区某某街道123号",
		},
		Sender: ManInfo{
			Name:    "李四",
			Mobile:  "13900139000",
			Address: "上海市浦东新区某某路456号",
		},
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
