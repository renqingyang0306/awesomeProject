package main

import (
	"fmt"
	"os"
)

func main() {
	//文件复制
	//file, err := os.Open("xiayang.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//file2, err := os.Create("xiayang2.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer file2.Close()
	//written, err := io.Copy(file2, file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(written)
	//
	// 目录操作
	dir, err := os.ReadDir("gostudy")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range dir {
		info, _ := v.Info()
		fmt.Println(v.Name(), "：", info.Size())
	}
}
