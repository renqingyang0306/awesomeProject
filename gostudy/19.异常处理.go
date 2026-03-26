package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

//func init() {
//	// 读取配置文件中，结果路径错了
//	_, err := os.ReadFile("xxx")
//	if err != nil {
//		// panic 会导致程序崩溃，应谨慎使用
//		panic(err.Error())
//	}
//}

func read() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 捕获异常，打印错误信息
			// 打印错误的堆栈信息
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) // 肯定会有一个panic
}

func Parent() error {
	err := method() // 遇到错误向上抛
	return err
}
func method() error {
	return errors.New("出错了")
}

func main() {
	//fmt.Println(Parent())

	read()
}
