package main

import (
	"fmt"
	"sort"
)

func arrayTest() {
	//一旦赋值或定义容量，那就是数组
	var nameList [3]string = [3]string{"shiomi", "kokoro"}
	//空判断，如果初始赋值了，那怕没东西，也是非空
	//false
	var nameList2 []string = []string{}
	fmt.Println(nameList2 == nil)
	//true
	var nameList3 []string
	fmt.Println(nameList3 == nil)
	//如果给予了空间，哪怕没赋值，无法进行空判断直接报错
	//invalid operation: nameList == nil (mismatched types [3]string and untyped nil)
	// fmt.Println(nameList == nil)
	fmt.Println(nameList)
	fmt.Println(nameList[1])
	//数组长度
	fmt.Println(len(nameList))

	//其他赋值方法
	var nameList4 = []string{}
	nameList5 := []string{}
	//make专门用于生成slice map channel内建引用类型，参数2表示切片的长度、图/映射的容量以及频道/通道的缓冲大小
	nameList6 := make([]string, 0)
	fmt.Println(nameList4, nameList5, nameList6)

	var userList []string
	userList = append(userList, "mimi")
	userList = append(userList, "nana")
	fmt.Println(userList)
	fmt.Println(userList[0])
	fmt.Println(cap(userList))

	//切片原理 就是将数组一个个切开，相当于其他语言中的slice方法
	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices)
	//数组切割，左开右闭
	fmt.Println(array[0:2])

	//排序
	ints := []int{4, 8, 2, 6}
	fmt.Printf("row: %v\n", ints)
	sort.Ints(ints)
	fmt.Printf("sort ASC: %v\n", ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Printf("sort DESC: %v\n", ints)
	//切片相关https://cloud.tencent.com/developer/article/2532107
}
