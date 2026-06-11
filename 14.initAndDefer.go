package main

import "fmt"

//初始化加载，顺序执行
//多个包，根据导入包的顺序执行
// var DB int

// func init() {
// 	DB = 10
// 	fmt.Println("init-1: db has loaded...")
// }
// func init() {
// 	fmt.Println("init-2")
// }
// func init() {
// 	fmt.Println("init-3")
// }

//注册延迟调用
//return前才会被调用，用来作资源清除
//等多个defer，先进后出（谁里return近线=先执行）
//尽管是return前执行，但如果想要拿到变量，也必须在defer函数前定义
func deferTest() int {
	var age int
	defer fmt.Println("defer-2")
	defer fmt.Println("defer-1")
	defer func() {
		fmt.Println("inDefer:", age)
	}()
	//报错
	// var age int
	return age
}
