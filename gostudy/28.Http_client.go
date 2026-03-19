package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 定义请求体结构体
type UserRequest struct {
	Username string `json:"username"`
	Data     string `json:"data"`
}

func main() {
	// http客户端
	// 实例化一个http客户端
	client := new(http.Client)
	// GET方式
	// 构造请求对象 - GET 请求
	//req, _ := http.NewRequest("GET", "http://localhost:8080/index", nil)
	//res, _ := client.Do(req)
	//// 获取响应
	//b, _ := io.ReadAll(res.Body)
	//fmt.Println(string(b))

	// 构造请求体数据
	requestBody := UserRequest{
		Username: "xiayang",
		Data:     "test data",
	}

	// 将结构体转换为JSON
	jsonData, _ := json.Marshal(requestBody)

	// 构造请求对象 - POST 请求，参数通过body传递
	req, _ := http.NewRequest("POST", "http://localhost:8080/index", bytes.NewBuffer(jsonData))
	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	// 发请求
	res, _ := client.Do(req)
	// 获取响应
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
}
