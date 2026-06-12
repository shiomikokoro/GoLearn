package main

import (
	"fmt"
	"time"
)

var doneChan1 = make(chan struct{})

func event1() {
	fmt.Println("eventFunction running.")
	time.Sleep(2 * time.Second)
	fmt.Println("eventFunction finished.")
	close(doneChan1)
}
func runEvent1() {
	go event1()
	select {
	case <-doneChan1:
		fmt.Println("Function Finished.")
	case <-time.After(1 * time.Second):
		fmt.Println("overtime")
		return
	}
}
