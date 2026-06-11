package main

import "fmt"

type bigNum struct {
	//存储数字
	digits [20]int
	//当前段数
	size int
}

func add(a bigNum, b bigNum) bigNum {
	var result bigNum
	var carry int = 0
	if a.size > b.size {
		result.size = a.size
	} else {
		result.size = b.size
	}
	for i := 0; i < result.size; i++ {
		var sum int = a.digits[i] + b.digits[i] + carry
		result.digits[i] = sum % 1e3
		fmt.Printf("运算次数%v，计算结果为%v \n", i+1, sum)
		carry = sum / 1e3
	}
	if carry != 0 {
		result.digits[result.size] = carry
		result.size++
	}
	return result

}
func formatNum(a bigNum) {
	for i := a.size - 1; i >= 0; i-- {
		fmt.Printf("%d", a.digits[i])
	}
	fmt.Printf("\n")
}
