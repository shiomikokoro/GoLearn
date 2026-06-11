package main

import (
	"encoding/json"
	"fmt"
)

// 结构体：类似于对象
type Class struct {
	Name string `json:"classname"`
	// className string
}

// tag
type Student struct {
	//组合class
	Class `json:"class"`
	Name  string   `json:"student_name"`
	Dream []string `json:"dreams"`
	//空值，提示不进行转换
	Age int `json:"age,omitempty"`
	//敏感信息，提示不进行转换
	Password string `json:"-"`
}

// 绑定函数
func (c Class) sayClassname() {
	fmt.Printf("classname is %s\n", c.Name)
}
func (s Student) study() {
	fmt.Printf("%s is studying\n", s.Name)
}
func (s Student) info() {
	//如果class中名字与student中的属性名不冲突，可以直接s.className
	fmt.Printf("ClassName: %s, StudentName: %s, Dream: %s\n", s.Class.Name, s.Name, s.Dream)
}

// 值传递并不能修改成功，需要引用地址
func (s *Student) SetName(name string) {
	s.Name = name
}

// 又因为进行引用传递了，如果传入的是slice/map类型的值，修改原本slice/map也会导致对象参数发生变化
func (s *Student) SetDream(dream []string) {
	s.Dream = make([]string, len(dream))
	copy(s.Dream, dream)
}
func runStudent() {
	s1 := Student{Name: "shiomi", Class: Class{Name: "999"}}
	s1.info()
	var name = "kokoro"
	s1.SetName(name)
	var dream = []string{"dream", "play game", "running"}
	s1.SetDream(dream)
	s1.info()
	dream[0] = "studying"
	s1.info()
	s1.sayClassname()

	//json标准化参数
	byteData, _ := json.Marshal(s1)
	fmt.Println(string(byteData))
}

//与继承的区别
//无法向上转型
//Person p = new Student(); //java没问题
//c:=Class{} s:=Student{} c=s //go报错，类型不一致
