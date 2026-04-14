package main

import (
	"errors"
	"fmt"
)

// 无返回值
func fun1() {
	return // 也可以不写
}

// 单返回值
func fun2() int {
	return 1
}

// 多返回值
func fun3() (int, error) {
	return 0, errors.New("错误")
}

// 命名返回值
func fun4() (res string) {
	return // 相当于先定义再赋值
	//return "abc"
}

func login() {
	fmt.Println("登录")
}

func userCenter() {
	fmt.Println("个人中心")
}

func logout() {
	fmt.Println("注销")
}

func main() {
	fmt.Println("请输入要执行的操作：")
	fmt.Println(`1：登录
			2：个人中心
			3：注销`)
	var num int
	fmt.Scan(&num)
	// 类型：map[int]func() 表示键的类型是 int，值的类型是 func()（即无参数、无返回值的函数）。
	var funcMap = map[int]func(){
		1: login,
		2: userCenter,
		3: logout,
	}
	funcMap[num]()
}
