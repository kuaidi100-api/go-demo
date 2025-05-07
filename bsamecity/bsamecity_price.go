package bsamecity

import (
	"encoding/json"
	"fmt"
	"time"
)

func Price() {
	method := "price"

	// 构建商品结构体
	goods := []Goods{
		{
			Name:  "外卖",
			Type:  "食品",
			Count: 0,
		},
	}

	// 构建请求参数
	param := Param{
		Kuaidicom:        "shunfengtongcheng",
		LbsType:          2,
		RecManName:       "张三",
		RecManMobile:     "13587654321",
		RecManProvince:   "北京市",
		RecManCity:       "北京市",
		RecManDistrict:   "海淀区",
		RecManAddr:       "学清嘉创大厦A座15层",
		RecManLat:        "40.014838",
		RecManLng:        "116.352569",
		SendManName:      "李四",
		SendManMobile:    "13512345678",
		SendManProvince:  "北京市",
		SendManCity:      "北京市",
		SendManDistrict:  "海淀区",
		SendManAddr:      "清华大学",
		SendManLat:       "40.002436",
		SendManLng:       "116.326582",
		Weight:           "1",
		Remark:           "测试下单",
		Volume:           "",
		OrderType:        0,
		ExpectPickupTime: "",
		ExpectFinishTime: "",
		Insurance:        "",
		Price:            "0",
		Goods:            goods,
		CallbackUrl:      "http://www.baidu.com",
		Salt:             "",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)

	// 发送请求
	_, err := DoBsamecityRequest(method, t, paramStr)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}
