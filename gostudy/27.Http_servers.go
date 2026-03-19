package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("请求地址：", r.URL.Path, r.UserAgent())
	w.Write([]byte("hello world"))
}

func handle(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		data, err := os.ReadFile("index.html")
		if err != nil {
			fmt.Println(data)
		}
		res.Write(data)
		//res.Write([]byte("<h1>hello 枫枫 GET</h1>"))
	case "POST":
		// 获取body数据
		data, err := io.ReadAll(req.Body)
		// 拿请求头
		contentType := req.Header.Get("Content-Type")
		fmt.Println(contentType)
		//switch contentType {
		//case "application/json":
		//
		//}

		if err != nil {
			fmt.Println(data)
		}
		ma := make(map[string]string)
		json.Unmarshal(data, &ma)
		fmt.Println(ma["username"])

		type User struct {
			Username string `json:"username"`
			Data     string `json:"data"`
		}
		var user User
		json.Unmarshal(data, &user)
		// 打印user json串
		fmt.Println(user.Username)
		userJson, _ := json.Marshal(user)
		fmt.Println(string(userJson))
		// 设置响应头
		header := res.Header()
		header["token"] = []string{"y1gyf156sdgT%d44hjgj"}
		res.Write([]byte("hello 枫枫 POST"))
		res.Write([]byte(user.Username))
	}
}

func main() {
	// 默认路由
	http.HandleFunc("/", index)
	// 根据地址路由
	http.Handle("/index", http.HandlerFunc(handle))
	// :8080  等价于 127.0.0.1:8080
	http.ListenAndServe(":8080", nil)
}
