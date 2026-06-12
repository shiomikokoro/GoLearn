package main

import (
	"fmt"
	"sync"
	"time"
)

//1GoroutineToXChannel

var moneyChan1 = make(chan int)
var nameChan = make(chan string)
var doneChan = make(chan struct{}) //用于关闭select信道

func pay1(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%v is paying.\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v was paid.\n", name)
	moneyChan1 <- money
	nameChan <- name
	wait.Done()
}

func startPay1() {
	var wait sync.WaitGroup

	startTime := time.Now()
	wait.Add(3)
	go pay1("shiomi", 1, &wait)
	go pay1("kokoro", 10, &wait)
	go pay1("misane", 2, &wait)

	go func() {
		defer close(moneyChan1)
		defer close(nameChan)
		defer close(doneChan)
		wait.Wait()
	}()
	var moneyList []int
	var nameList []string

	func() {
		for {
			select {
			case money := <-moneyChan1:
				moneyList = append(moneyList, money)
			case name := <-nameChan:
				nameList = append(nameList, name)
			case <-doneChan:
				return //终止函数
			}
		}
	}()
	fmt.Println("all pay finished, spead ", time.Since(startTime))
	fmt.Println("got ", moneyList, "dollors")
	fmt.Println(nameList, "come")
}
