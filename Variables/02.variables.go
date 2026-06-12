// 同文件夹下不能使用不同包名
package Variales

import "fmt"

//全局变量：可以不用
//不能用短声明符，只能在包内访问
var worldMessage = "message"

//多变量定义2：常见于全局或函数外定义
var (
	s1 string = "1"
	s2 string = "2"
)

//常量，必须声明了定义一起，之后不可修改
//如果数据想要跨包访问，首字母必须大写
const Version string = "0.0.1"

func init() {
	// fmt.Println("V-init-1")
}

//同文件夹/项目下只能有一个main入口
//如果数据想要跨包访问，方法名要大写
func Varibales() {
	//局部变量：声明了必须使用，不然编译不通过
	var name string
	name = "kaede"
	fmt.Printf(name)

	//声明赋值
	var name1 string = "kanade"
	fmt.Printf(name1)

	//省略类型
	var name2 = "kokoro"
	fmt.Printf(name2)

	//短声明符号
	name3 := "Wow"
	fmt.Printf(name3)

	//多变量定义1
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)

	test()

}
func test() {
	fmt.Println("该方法跨包无法使用，我实在包体被调用的")
}
