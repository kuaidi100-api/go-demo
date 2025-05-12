package work_order

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-test/config"
	"go-test/utils"
	"os"
	"strings"
	"time"
)

type WorkOrderUploadParam struct {
	FileName string `json:"fileName,omitempty"`
}

/*
*
上传附件
*/
func WorkOrderUpload(file *os.File) {

	param := WorkOrderUploadParam{
		FileName: file.Name(),
	}
	// 将参数转换为JSON字符串
	paramJson, _ := json.Marshal(param)
	paramStr := string(paramJson)

	// 生成时间戳
	t := fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	// 计算签名
	signStr := paramStr + t + config.KEY + config.SECRET
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	m := map[string]string{
		"key":   config.KEY,
		"sign":  sign,
		"t":     t,
		"param": paramStr,
	}

	// 发送请求
	_, err := utils.DoFileRequest(m, file, config.WORK_ORDER_UPLOAD_URL)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

}
