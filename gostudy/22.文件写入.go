package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 文件写入
	//const (
	//  O_RDONLY int = syscall.O_RDONLY // 只读
	//  O_WRONLY int = syscall.O_WRONLY // 只写
	//  O_RDWR   int = syscall.O_RDWR   // 读写
	//
	//  O_APPEND int = syscall.O_APPEND // 追加
	//  O_CREATE int = syscall.O_CREAT  // 如果不存在就创建
	//  O_EXCL   int = syscall.O_EXCL   // 文件必须不存在
	//  O_SYNC   int = syscall.O_SYNC   // 同步打开
	//  O_TRUNC  int = syscall.O_TRUNC  // 打开时清空文件
	//)
	//三个占位符
	//第一个是文件所有者所拥有的权限
	//第二个是文件所在组对其拥有的权限
	//第三个占位符是指其他人对文件拥有的权限
	//0444 表示三者均为只读的权限；
	//0666 表示三者均为“读写”的权限；
	//0777 表示三者均为读写执行的权限；
	//0764 表示所有者有读写执行（7=4+2+1）的权限，组有读写（6=4+2）的权限，其他用户则为只读（4=4）；
	file, err := os.OpenFile("xiayang1.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 如果文件存在先清空
	file.Truncate(0)
	defer file.Close()
	for i := 0; i < 1000; i++ {
		file.WriteString("测试" + strconv.Itoa(i) + "\n")
	}
}
