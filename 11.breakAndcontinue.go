package main

import "fmt"

func kukuno() {
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d*%d=%d \t", i, j, i*j)
	// 	}
	// 	fmt.Printf("\n")
	// }

	//break 中断循环
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				break
			}
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Printf("\n")
	}
}
func continueTest() {
	for i := 0; i < 10; i++ {
		//continue 跳过本轮
		if i == 5 {
			continue
		}
		fmt.Println(i)

	}
}
