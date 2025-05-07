package base

import (
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
)

type MapTrackParam struct {
	Com      string `json:"com"`
	Num      string `json:"num"`
	Phone    string `json:"phone"`
	From     string `json:"from"`
	To       string `json:"to"`
	Resultv2 string `json:"resultv2"`
	Show     string `json:"show"`
	Order    string `json:"order"`
}

/*
快递查询地图轨迹
*/
func MapTrack() {

	param := MapTrackParam{
		Com:      "ems",
		Num:      "em263999513jp",
		Phone:    "13868688888",
		From:     "广东省深圳市南山区金蝶软件园",
		To:       "北京朝阳区国际金融大厦",
		Resultv2: "5",
		Show:     "0",
		Order:    "desc",
	}

	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 发送请求
	_, err := utils.CustomerRequest(paramStr, config.MAP_TRACK_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
