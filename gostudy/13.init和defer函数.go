package main

import "fmt"

func init() {
	fmt.Println("init1")
}
func init() {
	fmt.Println("init2")
}

func Func() {
	defer fmt.Println("defer2")
	fmt.Println("func")
	defer fmt.Println("defer1")
}

func main() {
	// 多个defer语句，按先进后出的方式执行，谁离return近谁先执行
	defer fmt.Println("defer4")
	Func()
	defer fmt.Println("defer3")
}
