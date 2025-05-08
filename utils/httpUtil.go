package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go-test/config"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func DoRequest(t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + config.KEY + config.SECRET
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", config.KEY)
	formData.Add("t", t)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return execute(postUrl, formData)
}

/*
*
多一个Method入参
*/
func DoMethodRequest(method string, t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + config.KEY + config.SECRET
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", config.KEY)
	formData.Add("method", method)
	formData.Add("t", t)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return execute(postUrl, formData)
}

/*
*
使用customer鉴权
*/
func CustomerRequest(param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + config.KEY + config.CUSTOMER
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("customer", config.CUSTOMER)
	formData.Add("sign", sign)
	formData.Add("param", param)

	return execute(postUrl, formData)
}

/*
*
根据map传入form数据
*/
func DoMapRequest(m map[string]string, postUrl string) (string, error) {
	formData := url.Values{}
	// 由map生成form表单数据
	for key, value := range m {
		formData.Add(key, value)
	}
	return execute(postUrl, formData)
}

/**
*执行HTTP请求
 */
func execute(postUrl string, formData url.Values) (string, error) {
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
