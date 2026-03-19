package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 一次性读取
	//file, err := os.ReadFile("xiayang.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(file))
	//带缓冲区读取
	// 分片读取
	//open, err := os.Open("xiayang.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer open.Close()
	//for {
	//	buf := make([]byte, 1024)
	//	n, err := open.Read(buf)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//	}
	//	fmt.Println(string(buf[:n]))
	//}
	// 带缓冲区的按行读取
	open, err := os.Open("xiayang.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer open.Close()
	//reader := bufio.NewReader(open)
	//for {
	//	line, _, err := reader.ReadLine()
	//	fmt.Println(string(line))
	//	if err != nil {
	//		break
	//	}
	//}
	// 按照指定分割符
	scanner := bufio.NewScanner(open)
	scanner.Split(bufio.ScanWords) // 按照单词读
	//scanner.Split(bufio.ScanLines) // 按照行读
	//scanner.Split(bufio.ScanRunes) // 按照中文字符读
	//scanner.Split(bufio.ScanBytes) // 按照字节读读，中文会乱码
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
