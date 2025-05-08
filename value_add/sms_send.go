package value_add

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"strings"
)

/*
*快递100短信发送接口
 */
func SmsSend() {

	signStr := config.KEY + config.USER_ID
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	m := map[string]string{
		"sign":     sign,
		"userid":   "9974ef2c377a4dbt9",
		"seller":   "快递100",
		"phone":    "13568688888",
		"tid":      "11",
		"content":  "{\"接收人姓名\":\"王帅\", \"公司名\":\"快递100\", \"快递单号\":\"154893238584\", \"url\":https: //api.kuaidi100.com/home\"}",
		"outorder": "143255893",
		"callback": "http://xxx/callback",
	}

	// 发送请求
	_, err := utils.DoMapRequest(m, config.SMS_SEND_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
