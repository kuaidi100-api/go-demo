package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	KEY    = "***" // 客户授权key
	SECRET = "***" // 秘钥
)

type Payload struct {
	Key    string `json:"key"`
	Method string `json:"method"`
	T      string `json:"t"`
	Param  string `json:"param"`
	Sign   string `json:"sign"`
}

func DoRequest(method string, t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + KEY + SECRET
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构建请求体
	payload := Payload{
		Key:    KEY,
		Method: method,
		T:      t,
		Param:  param,
		Sign:   sign,
	}

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", payload.Key)
	formData.Add("method", payload.Method)
	formData.Add("t", payload.T)
	formData.Add("sign", payload.Sign)
	formData.Add("param", payload.Param)

	// 创建HTTP客户端
	client := &http.Client{}

	// 创建HTTP请求
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(formData.Encode()))
	if err != nil {
		fmt.Printf("创建请求失败: %v\n", err)
		return "创建请求失败", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 发送HTTP请求
	fmt.Printf("请求信息: %v\n", req)
	fmt.Printf("请求参数: %v\n", formData)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return "请求失败", err
	}
	defer resp.Body.Close()

	// 读取响应内容
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return string(respBody), err
	}

	// 打印响应内容
	fmt.Println("响应内容:", string(respBody))
	return string(respBody), err
}
