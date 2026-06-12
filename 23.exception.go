package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

// 向上抛，实际上就是返回类型为error
func Servise1() (res int, err error) {
	res, err = method1(1, 0)
	if err != nil {
		return
	}
	//进行其他数据处理

	return
}
func method1(a int, b int) (num int, err error) {
	if b == 0 {
		err = errors.New("除0异常")
		return
	}
	num = a / b
	return
}

// 直接中断 该类一般放在init()函数中，因为后续程序启动已经不太重要了
func readFile1() {
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}

// 恢复程序/保持运行
// 用defer去接无法获得的错误（返回值不带err），来实现对panic的捕获，保证不会使程序崩溃
// 一般在框架层的异常处理所做的
func read1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)           //因为这么做不知道哪儿错了
			s := string(debug.Stack()) //所以调用堆栈信息还原报错
			fmt.Println(s)             //继续运行此后的代码
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2])            //越界异常
	fmt.Println("Program running.") //异常后无法运行
}
