package main

import (
	"awesomeProject/renqy"
	"fmt"
	"sort"
)

/*
*
TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
*/
func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	fmt.Println("")

	fmt.Println(renqy.Auth)

	var z rune = '你'
	fmt.Printf("%c %d\n", z, z)

	var list = []int{6, 2, 3}
	fmt.Println("排序前：", list)
	sort.Ints(list)
	fmt.Println("排序后：", list)
	// 逆序
	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	fmt.Println("排序后2：", list)

	var name = make([]string, 3)
	fmt.Println(name, len(name), cap(name))
	fmt.Printf("%#v \n", name[1])
	fmt.Println(name == nil) // false
}
