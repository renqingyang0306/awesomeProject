package main

import (
	"fmt"
	"sync"
	"time"
)

var num int

// sync.WaitGroup 用于等待多个 Goroutine 完成。
var wait sync.WaitGroup

// 同步锁
var lock sync.Mutex

// 定义一个map
// var mp = make(map[string]string)
// 线程安全map
var mp = sync.Map{}

func add() {
	// 谁先抢到了这把锁，谁就把它锁上，一旦锁上，其他的线程就只能等着
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num++
	}
	lock.Unlock()
	wait.Done()
}
func reduce() {
	lock.Lock()
	for i := 0; i < 1000000; i++ {
		num--
	}
	lock.Unlock()
	wait.Done()
}

//	func reader() {
//		for {
//			lock.Lock()
//			fmt.Println(mp["time"])
//			lock.Unlock()
//		}
//		wait.Done()
//	}
//
//	func writer() {
//		for {
//			//lock.Lock()
//			mp["time"] = time.Now().Format("15:04:05")
//			//lock.Unlock()
//		}
//		wait.Done()
//	}
func reader() {
	for {
		value, _ := mp.Load("time")
		fmt.Println(value)
	}
	wait.Done()
}
func writer() {
	for {
		mp.Store("time", time.Now().Format("15:04:05"))
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	//go add()
	//go reduce()
	//fmt.Println(num)
	//我们不能在并发模式下读写map
	//如果要这样做
	//给读写操作加锁
	//使用sync.Map
	go reader()
	go writer()
	wait.Wait()

}
