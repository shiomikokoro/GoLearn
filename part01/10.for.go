package main

import "fmt"

func forTest() {
	var num int = 0
	var list = []string{"shiomi", "kokoro"}
	var userMap = map[string]string{"name": "shiomi", "age": "0"}
	for i := 0; i < 100; i++ {
		num++
	}
	fmt.Println(num)
	for index, item := range list {
		fmt.Println("list:", index, item)
	}
	for key, value := range userMap {
		fmt.Println("map:", key, value)
	}
}
