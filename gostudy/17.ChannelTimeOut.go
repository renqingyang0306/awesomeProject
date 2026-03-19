package main

import (
	"fmt"
	"time"
)

// var done = make(chan struct{})
//
//	func event() {
//		fmt.Println("event执行开始")
//		time.Sleep(2 * time.Second)
//		fmt.Println("event执行结束")
//		close(done)
//	}
//
//	func main() {
//		go event()
//
//		select {
//		case <-done:
//			fmt.Println("协程执行完毕")
//		case <-time.After(1 * time.Second):
//			fmt.Println("超时")
//			return
//		}
//	}
func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		//time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		//time.Sleep(5 * time.Second)
	}()
	time.Sleep(1 * time.Second)
	//select中各个case执行顺序是随机的；
	//如果某个case中的channel已经ready，则执行相应的语句并退出select流程，
	//如果所有case中的channel都未ready，则执行default中的语句然后退出select流程；
	select {
	case <-chan1:
		fmt.Println("chan1 ready.")
	case <-chan2:
		fmt.Println("chan2 ready.")
	default:
		fmt.Println("default")
	}

	fmt.Println("main exit.")
}
