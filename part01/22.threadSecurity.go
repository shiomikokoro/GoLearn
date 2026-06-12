package main

import (
	"fmt"
	"sync"
)

var sum1 int
var wait sync.WaitGroup

func add2() {
	for i := 0; i < 100000; i++ {
		sum1++
	}
	wait.Done()
}
func sub2() {
	for i := 0; i < 100000; i++ {
		sum1--
	}
	wait.Done()
}

// 一般而言，运行出来的结果不一定为0
// 为0，说明cpu还不错，但一定要处理这类“并发问题”
func runProgram1() {
	wait.Add(2)
	go add2()
	go sub2()
	fmt.Println(sum1)
}

// 所以对数据进行加减/数据操作的过程加锁
var lock1 sync.Mutex

func add3() {
	lock1.Lock()
	for i := 0; i < 100000; i++ {
		sum1++
	}
	lock1.Unlock()
	wait.Done()
}
func sub3() {
	lock1.Lock()
	for i := 0; i < 100000; i++ {
		sum1--
	}
	lock1.Unlock()
	wait.Done()
}
func runProgram2() {
	wait.Add(2)
	go add3()
	go sub3()
	fmt.Println(sum1)
}

// map类数据也会因为并发读写导致线程不安全
// 会报错强制停止，concurrent map read and map write
var map1 = map[int]string{}

func runProgram3() {
	go func() {
		for {
			map1[1] = "shiomi"
		}
	}()
	go func() {
		for {
			fmt.Println(map1[1])
		}
	}()
	select {} //阻断结束看效果
}

// 所以使用sync的map类进行高并发
var map2 = sync.Map{}

func runProgram4() {
	go func() {
		for {
			map2.Store(1, "shiomi")
		}
	}()
	go func() {
		for {
			fmt.Println(map2.Load(1))
		}
	}()
	select {} //阻断结束看效果
}
