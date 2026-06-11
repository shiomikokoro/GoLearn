package main

import (
	"fmt"
	"sync"
	"time"
)

// channel 通道
// 用于解决需要全局变量接受协程的返回值
// 需要初始化，不过长度为0
var moneyChan = make(chan int)

func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%v is paying.\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v was paid.\n", name)
	moneyChan <- money
	wait.Done()
}

func startPay() {
	var wait sync.WaitGroup

	startTime := time.Now()
	wait.Add(3)
	go pay("shiomi", 1, &wait)
	go pay("kokoro", 10, &wait)
	go pay("misane", 2, &wait)
	//在等待前获取值，避免死锁
	//利用协程监听
	go func() {
		defer close(moneyChan)
		wait.Wait() //没有数据了关闭监听协程
	}()

	var moneyList []int
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}
	wait.Wait()
	fmt.Println("all pay finished, spead ", time.Since(startTime))
	fmt.Println("got ", moneyList, "dollors")
}
