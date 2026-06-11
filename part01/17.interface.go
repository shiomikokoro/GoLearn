package main

import "fmt"

type Chicken struct {
	Name string
}
type Cat struct {
	Name string
}

//这么做既麻烦，又不能同名，还会报错
// func (c Chicken) sing() {
// 	fmt.Printf("%v is sing", c.Name)
// }
// func (c Cat) sing() {
// 	fmt.Printf("%v is sing", c.Name)
// }
// func sing(c Chicken) {
// 	c.sing()
// }
// func sing(c Cat) {//同名报错
// 	c.sing()
// }

type SingInterface interface {
	sing() //同名
}

//实现接口
func (c Chicken) sing() { //同名
	fmt.Printf("%v is sing\n", c.Name)
}
func (c Cat) sing() { //同名
	fmt.Printf("%v is sing\n", c.Name)
}
func Sing(c SingInterface) { //参数为接口
	//接口断言 c.(Chicken)
	// ch, ok := c.(Chicken)
	// fmt.Println(ch, ok)
	//一般采取以下方法 switch可以拿到具体赋值
	switch class := c.(type) {
	case Chicken:
		fmt.Printf("%T \n", class)
	case Cat:
		fmt.Printf("%T \n", class)
	default:
		fmt.Println("unknow")

	}
	c.sing()
}

//空接口表示任何类型都实现了该定义
type EmptyInterface interface{}

func Print(val EmptyInterface) {
	fmt.Println(val)
}

//更简洁的写法
// func Print(val interface{}){}
// func Print(val any){}//interface{}的别名

func chorus() {
	cat := Cat{Name: "cat"}
	chicken := Chicken{Name: "ik"}
	Sing(cat)
	Sing(chicken)
	Print(0)
	Print(cat)
}
