package main

import (
	"encoding/json"
	"fmt"
)

// 泛型
func add4(n1, n2 int) int {
	return n1 + n2
}

// 如果没有泛型需要重新写一个方法
func add5(n1, n2 uint) uint {
	return n1 + n2
}
func runIntAdd() {
	add4(1, 2)
	add5(1, 2)
	//或者强转
	var u1, u2 = uint(1), uint(2)
	add4(int(u1), int(u2))
}

// 现在可以用T来表示泛型 2022支持的
// 但要记住T只能表达一个类型，函数需要通过[T]来理解有哪些类型
func add6[T int | uint](n1, n2 T) T {
	return n1 + n2
}
func runAdd() {
	add6(1, 2)
	add6(uint(1), uint(2))
}

// 多类型可以用K表示
func myPrint[T int, K string | int](n1 T, n2 K) {}

// 当然加法数字类型太多，可以用接口来解决
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func add7[T Number](n1, n2 T) T {
	return n1 + n2
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}
type User struct {
	Name string `json:"name"`
}
type Userinfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUserJson() {
	res1 := Response{
		Code: 1,
		Msg:  "success",
		Data: User{
			Name: "shiomi",
		},
	}
	byteData1, _ := json.Marshal(res1)
	res2 := Response{
		Code: 1,
		Msg:  "success",
		Data: Userinfo{
			Name: "shiomi",
			Age:  1,
		},
	}
	byteData2, _ := json.Marshal(res2)
	fmt.Printf("User:%v\n", string(byteData1))
	fmt.Printf("User:%v\n", string(byteData2))
}

// 泛型反解
type Response2[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func unmarshal() {
	var userResponse Response
	var userResponse2 Response2[User]
	var userResponse3 Response2[Userinfo]
	json.Unmarshal([]byte(`{"code":1,"msg":"success","data":{"name":"shiomi"}}`), &userResponse)
	json.Unmarshal([]byte(`{"code":1,"msg":"success","data":{"name":"shiomi"}}`), &userResponse2)
	fmt.Println(userResponse)  //变成了map
	fmt.Println(userResponse2) //正常连点获取数据
	json.Unmarshal([]byte(`{"code":1,"msg":"success","data":{"name":"shiomi","age":1}}`), &userResponse3)
	fmt.Println(userResponse3.Data.Age)

}

// 泛型切片
type MySlice[T any] []T

func genericeSlice() {
	var stringSlice MySlice[string]
	stringSlice = append(stringSlice, "shiomi")
	var intSlice MySlice[int]
	intSlice = append(intSlice, 1)
}

// 泛型map
// key 必须为简单（基本）类型
type MyMap[K string | int, V any] map[K]V

func genericeMap() {
	var stringMap = make(MyMap[string, string])
	stringMap["name"] = "shiomi"
	var intMap = make(MyMap[int, int])
	intMap[1] = 2
}
