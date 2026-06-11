package main

import "fmt"

//计算机运行原理
//所有数据都是在内存里面的，拉去硬盘数据后在内存分配一个空间

//两个name不一样，互不影响
func copyName(name string) {
	fmt.Printf("inFun: %p \n", &name)
	name = "kokoro"
}
func originName() {
	var name = "shiomi"
	fmt.Printf("mainFun: %p \n", &name)
	copyName(name)
	fmt.Printf("afterFuncChange: %v \n", name)
	pointCopy(&name)
	fmt.Printf("afterFuncChange: %v \n", name)
}

//引用传递，传递的参数是地址
func pointCopy(name *string) {
	fmt.Printf("inPoint: %p \n", name)
	*name = "misane"
}
