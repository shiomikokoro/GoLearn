package main

import (
	"fmt"
	"io"
	"os"
)

const (
	O_RDONLY string = "O_RDONLY" // 只读
	O_WRONLY string = "O_WRONLY" // 只写
	O_RDWR   string = "O_RDWR  " // 读写

	O_APPEND string = "O_APPEND" // 追加
	O_CREATE string = "O_CREAT " // 如果不存在就创建
	O_EXCL   string = "O_EXCL  " // 文件必须不存在
	O_SYNC   string = "O_SYNC  " // 同步打开
	O_TRUNC  string = "O_TRUNC " // 打开时清空文件
)

func openFile() {
	//参数filepath flag perm
	//flag 文件打开方式
	//perm linux系统中文件的可执行权限
	//R Read 八进制 4
	//W Wirte 八进制 2
	//X execute 八进制 1
	//三个数字分别代表所有者、组、其他用户的权限
	//0444 三者都只读
	//0666 三者都能读写
	//0777 三者都能读写执行
	//0764 所有者读写执行、组能读写、其他用户只读
	file, err := os.OpenFile("hello.log", os.O_RDWR|os.O_CREATE, 0764)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	byteData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteData))

}

// 最简单文件写入
func fileWirte() {
	err := os.WriteFile("hello.log", []byte("你好"), 0764)
	fmt.Println(err)
}

// 文件复制
func fileCopy() {
	read, _ := os.Open("file1.txt")
	defer read.Close()
	write, _ := os.Create("file3.txt") // 默认是 可读可写，不存在就创建，清空文件
	defer write.Close()
	n, err := io.Copy(write, read)
	fmt.Println(n, err)
}

// 目录控制
func dirController() {
	dir, _ := os.ReadDir(".")
	for _, entry := range dir {
		info, _ := entry.Info()
		fmt.Println(entry.Name(), info.Size()) // 文件名，文件大小，单位比特
	}
}
