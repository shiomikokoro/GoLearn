package main

import "fmt"

//输出
func outputTest() {
	var name = "shiomi"
	//带换行，可输入多个值，自带空格
	fmt.Println("hello", "world")
	fmt.Println("by shiomi")
	//格式化打印
	fmt.Printf("string:hello,%s \n", "shiomi")
	fmt.Printf("class:hello,%T \n", "shiomi")
	fmt.Printf("int:hello,%d \n", 1)
	//打印 1.00 数字几保留几位小数，不写全保留
	fmt.Printf("float:hello,%.2f \n", 1.0)
	fmt.Printf("float:hello,%c \n", 'a')
	fmt.Printf("anything:hello,%v \n", 1.0)
	//打印 "1" 用于打印空字符串常见
	fmt.Printf("fullFormat:hello,%#v \n", "1")
	//中
	fmt.Printf("rune: %c", 20013)
	//指针、内存地址
	fmt.Printf("point:hello,%p \n", &name)

	//格式化赋值
	var trans = fmt.Sprintf("%.2f", 3.1415)
	fmt.Println(trans)
}

//输入，一般用于控制台系统
func inputTest() {
	fmt.Print("输入您的名字：")
	var name string
	//指针
	fmt.Scan(&name)
	fmt.Printf("您刚才输入的名字是：%v \n", name)
	fmt.Print("输入您的年龄：")
	var age int
	n, err := fmt.Scan(&age)
	//正常 age 1 <nil>（其它语言中的null） 错误 0 0 expected integer
	//int未赋值默认值为0
	fmt.Printf("您刚才输入的年龄是：%v，有效码：%v，错误信息：%v", age, n, err)
}
