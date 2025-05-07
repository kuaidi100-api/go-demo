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

type Payload struct {
	Key      string `json:"key,omitempty"`
	Method   string `json:"method,omitempty"`
	T        string `json:"t,omitempty"`
	Param    string `json:"param,omitempty"`
	Sign     string `json:"sign,omitempty"`
	Customer string `json:"customer,omitempty"`
}

func DoRequest(t string, param string, postUrl string) (string, error) {
	// 计算签名
	signStr := param + t + config.KEY + config.SECRET
	hash := md5.New()
	hash.Write([]byte(signStr))
	sign := hex.EncodeToString(hash.Sum(nil))
	sign = strings.ToUpper(sign)

	// 构建请求体
	payload := Payload{
		Key:   config.KEY,
		T:     t,
		Param: param,
		Sign:  sign,
	}

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("key", payload.Key)
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

	// 构建请求体
	payload := Payload{
		Key:    config.KEY,
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

	// 构建请求体
	payload := Payload{
		Customer: config.CUSTOMER,
		Param:    param,
		Sign:     sign,
	}

	// 构造form表单数据
	formData := url.Values{}
	formData.Add("customer", payload.Customer)
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
