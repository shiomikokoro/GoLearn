package main

import "fmt"

// 如果直接使用int表示，会导致代码部分不够直观
// 且可以接受其他int数据
// const (
// 	SuccessCode    = 0
// 	ServiceErrCode = 1001
// 	NetworkErrCode = 1002
// )
//所以打包成自定义类型
type Code int

const (
	SuccessCode    Code = 0
	ServiceErrCode Code = 1001
	NetworkErrCode Code = 1002
)

func (c Code) GetMsg() string {
	switch c {
	case SuccessCode:
		return "Success"
	case ServiceErrCode:
		return "ServiceError"
	case NetworkErrCode:
		return "NetworkError"
	default:
		return ""
	}
}
func webService(name string) (code Code, msg string) {
	if name == "1" {
		// 这样就不需要创建函数去获取，可以直接使用绑定的函数，还不会识别别的数据
		// return ServiceErrCode, GetMsg(1)
		return ServiceErrCode, ServiceErrCode.GetMsg()
	} else if name == "2" {
		return NetworkErrCode, NetworkErrCode.GetMsg()
	} else {
		return SuccessCode, SuccessCode.GetMsg()
	}
}

//类型别名是不能绑定方法的
//别名，这个表示int的别名是AlisaCode
type AlisaCode = int

// func (c AlisaCode) GetMsg() {} //报错

const myCode Code = 1
const myAlisaCode AlisaCode = 1

func AlisaTest() {
	//打印类型还是原始类型
	fmt.Printf("%v,%T\n", myCode, myCode)           //myCode
	fmt.Printf("%v,%T\n", myAlisaCode, myAlisaCode) //int
	//原始类型和包装类型不能做断言判断
	if myCode == 1 {
	}
	// if myCode == myAlisaCode{}//报错 myCode是Code类，myAlisaCode是int类
}
