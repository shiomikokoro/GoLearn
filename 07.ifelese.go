package main

import "fmt"

func ifelse() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	// 中断时，卫语句
	// if age<=0 {
	// 	fmt.Printf("未出生")
	// 	return
	// }

	if age <= 0 {
		fmt.Printf("未出生")
	} else if age <= 18 {
		fmt.Printf("未成年")
	} else if age <= 35 {
		fmt.Printf("青年")
	} else {
		fmt.Printf("中年")
	}
	//不建议ifelse内嵌套ifelse
	//逻辑符号不再重复啰嗦 && || !
}
