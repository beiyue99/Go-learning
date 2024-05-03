package main

import (
	"fmt"
	"time"
)

func newTask() {
	for i := 0; ; i++ {
		fmt.Printf("new Goroutine :i=%d\n", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask() //创建协程
	i := 0
	for {
		i++
		fmt.Printf("main Goroutine:i=%d\n", i)
		time.Sleep(time.Second)
	}
}
