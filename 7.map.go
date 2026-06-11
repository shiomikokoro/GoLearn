package main

import "fmt"

//字典、映射、图
//创建时必须初始化 jason?
func mapTest() {
	var userMap map[int]string = map[int]string{
		1: "shiomi",
		2: "kokoro",
		4: "",
	}
	fmt.Println(userMap)
	//类似数字方式，不过参数是[key]
	fmt.Println(userMap[1])
	fmt.Printf("%#v\n", userMap[3])
	//获得对应[key]的内容
	value1 := userMap[3]
	fmt.Printf("%#v\n", value1)
	//获得对应[key]的内容，flag用于接受是否为空内容
	value2, flag := userMap[3]
	fmt.Printf("%#v\n,%v", value2, flag)

	//CURD
	userMap[1] = "misaki"
	userMap[3] = "murasaki"
	fmt.Println(userMap)
	delete(userMap, 4)
	fmt.Println(userMap)

	//创建时必须初始化
	//ERR panic: assignment to entry in nil map
	// var map1 map[string]string
	// map1["name"] = "shiomi"
	//空初始化
	var map1 = map[string]string{}
	map1["name"] = "shiomi"
	var map2 = make(map[string]string, 0)
	map2["name"] = "kokoro"
	fmt.Println(map2)

}
