package main

import "fmt"

func switchTest() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)

	//和其他语言不一样，默认不需要break
	//可以使用判断条件
	switch {
	case age <= 0:
		fmt.Printf("未出生")
	case age <= 18:
		fmt.Printf("未成年")
		//如果想要继续进行判断使用 fallthrough
		fallthrough
	case age <= 35:
		fmt.Printf("青年")
	default:
		fmt.Printf("中年")
	}
	var week int
	fmt.Println("请输入你的日期：")
	fmt.Scan(&week)

	switch week {
	case 1:
		fmt.Printf("周1")
	case 2:
		fmt.Printf("周2")
	case 3:
		fmt.Printf("周3")
	case 4:
		fmt.Printf("周4")
	case 5:
		fmt.Printf("周5")
	case 6, 7:
		fmt.Printf("休息日")
	default:
		fmt.Printf("非法数据")

	}
}
