package part02__test

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"
)

// 向上抛，实际上就是返回类型为error
func Parent1() error {
	err := method1()
	return err
}
func method1() error {
	return errors.New("出错了")
}

// 直接中断 该类一般放在init()函数中，因为后续程序启动已经不太重要了
func readFile1() {
	_, err := os.ReadFile("xxx")
	if err != nil {
		panic(err.Error())
	}
}

// 恢复程序/保持运行
// 用defer去接，来实现对panic的捕获，保证不会使程序崩溃
// 一般在框架层的异常处理所做的
func read1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			s := string(debug.Stack())
			fmt.Println(s)
		}
	}()
	var list = []int{2, 3}
	fmt.Println(list[2]) //越界异常
}
