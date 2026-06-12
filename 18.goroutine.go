package main

import (
	"fmt"
	"sync"
	"time"
)

// 全局信号量
// var wait int

// 协程
//
//	func shopping(name string) {
//		fmt.Printf("%v is shopping.\n", name)
//		time.Sleep(1 * time.Second)
//		fmt.Printf("%v was shopped.\n", name)
//		wait.Done()
//	}
func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%v is shopping.\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v was shopped.\n", name)
	wait.Done()
}

func startShopping() {
	// 官方设置的等待结构体
	var wait sync.WaitGroup
	startTime := time.Now()
	// wait = 3
	// 有几个需要等待的并发
	wait.Add(3)
	// 顺序购物
	// shopping("shiomi")
	// shopping("kokoro")
	// shopping("misane")
	// 并发（一起去）购物
	// go shopping("shiomi")
	// go shopping("kokoro")
	// go shopping("misane")
	// 传入指针
	go shopping("shiomi", &wait)
	go shopping("kokoro", &wait)
	go shopping("misane", &wait)
	// 不过别人还在购物，主线程还在运行，主进程结束会导致并发一起结束
	// 如果使用time.Sleep()强制等待，会导致并发还没结束或者等待太久了
	// 所以在并行之前设置一个信号量，表示进程运行情况
	// for {
	// 	if wait == 0 {
	// 		break
	// 	}
	// }
	// 不过官方有相关设置
	// 这样的好处就不需要设置全局变量了
	wait.Wait()
	fmt.Println("all shopping finished, spead ", time.Since(startTime))
}
