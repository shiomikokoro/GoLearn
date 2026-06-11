package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("hello")
}

// 传参
func addFunc(a int, b int) {
	fmt.Printf("%d + %d = %d", a, b, a+b)
}

// 超多参数
func addMoreArgu(numberList ...int) {
	var sum int
	for _, i := range numberList {
		sum += i
	}
	fmt.Println(sum)

}

// 返回值
// 空返回也可以，用于终止函数运行
func addHasReturnFunc(a int, b int) int {
	return a + b
}

// 多返回值1
func returnSomething() (string, bool) {
	if 1 > 2 {
		return "", false
	}
	if 1 > 3 {
		return "", false
	}
	return "", false
}

// 精简化多返回值，return直接返回对应参数返回内容
// 中途可以变更
// 如果没有变更，就是各数据类型的零值
func returnSomething2() (val string, flag bool) {
	if 1 > 2 {
		val = "12"
		return
	}
	return
}

// 匿名函数
// go禁止在函数内在定义函数
// 但可以用匿名函数
func outFunc() {
	var inFunc = func(name string) {
		fmt.Println(name)
	}
	inFunc("name is guess")
}

// 高级函数，将函数作为参数传递
func advancedFunc() {
	fmt.Println("请输入你想执行的操作")
	fmt.Println("1. 登录")
	fmt.Println("2. 注册")
	var index int
	fmt.Scan(&index)

	var funMap = map[int]func(){
		1: login,
		2: register,
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	}
	// switch index {
	// case 1:
	// 	login()
	// case 2:
	// 	register()
	// default:
	// 	return
	// }

}
func login()    {}
func register() {}

// 延迟两秒执行1+2+3
func closureAdd(awaitSecond int) func(...int) int {
	// 一般在外面使用，不算闭包
	// 直接调用closureAdd(awaitSecond)会返回函数，再调用这个函数时则不会延迟执行
	// time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(numberList ...int) (sum int) {
		time.Sleep(time.Duration(awaitSecond) * time.Second)
		for _, v := range numberList {
			sum += v
		}
		return sum
	}
}

func runClosureAdd(awaitSecond int, numberList ...int) {
	t1 := time.Now()
	sum := closureAdd(awaitSecond)(numberList...)
	subTime := time.Since(t1)
	fmt.Println("add:", sum, "runTime:", subTime)
}
