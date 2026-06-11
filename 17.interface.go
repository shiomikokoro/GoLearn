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
	c.sing()
}

func chorus() {
	cat := Cat{Name: "cat"}
	chicken := Chicken{Name: "ik"}
	Sing(cat)
	Sing(chicken)
}
