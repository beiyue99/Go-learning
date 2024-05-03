package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("gorouline结束！")
		fmt.Println("gorouline 正在运行！")

		c <- 556
	}()
	num := <-c
	fmt.Println("num=", num)
	fmt.Println("main gorouline 结束!")
}

//为什么主程永远可以正确打印num的值？因为channel会阻塞两边
//如果c还没有数据，就会阻塞等待有数据，然后才被num获取
