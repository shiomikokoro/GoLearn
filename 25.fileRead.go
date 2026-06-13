package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func fileRead() {
	byteData, err := os.ReadFile("hello.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteData))
}

// 分片读取
func sliceFileRead() {
	file, err := os.Open("hello.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var byteData = make([]byte, 13) //分片
	for {
		n, err := file.Read(byteData)
		if err == io.EOF {
			break
		}
		fmt.Printf("content:%#v,size:%d \n", string(byteData), n)
	}
}

// 分行读取
func rowFileRead() {
	file, err := os.Open("hello.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}

// 分隔符读取
func sepFileRead() {
	file, err := os.Open("hello.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewScanner(file)
	buf.Split(bufio.ScanLines) //分割规则
	var index int
	for buf.Scan() {
		index++
		fmt.Println(buf.Text(), " ", index)
	}
	if err := buf.Err(); err != nil { //查看扫描是否出错，可能未扫完发生异常
		fmt.Printf("scan error: %v", err) // 处理错误，例如记录日志、返回错误等
	}
}
