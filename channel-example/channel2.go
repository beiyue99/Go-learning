package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //最大容量是3
	fmt.Println("len(c)=", len(c), "cap(c)=", cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行，发送的元素=", i, "len(c)=", len(c), "cap(c)=", cap(c))

		}
	}()

	time.Sleep(2 * time.Second)
	//子go先发了3个，然后阻塞等待主go取

	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println("num=", num)
	}
	fmt.Println("main 结束！")
}
