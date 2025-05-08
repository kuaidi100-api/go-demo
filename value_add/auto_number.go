package value_add

import (
	"fmt"
	"go-test/config"
	"go-test/utils"
)

/*
*
快递智能识别单号
*/
func AutoNumber() {

	number := "906919164534"
	key := config.KEY
	postUrl := config.AUTO_NUMBER_URL + "?" + "num=" + number + "&key=" + key
	m := map[string]string{}

	// 发送请求
	_, err := utils.DoMapRequest(m, postUrl)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
