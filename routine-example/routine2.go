package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			//如果在这里写return，那么不会打印B，但会打印A
			//如果用协程退出函数则都会退出，不打印AB
			//runtime.Goexit()
			fmt.Println("B")
		}() //小括号代表调用匿名函数
		fmt.Println("A")
	}()

	//再创建一个有参数和返回值的协程函数
	//由于返回的值不在当前主进程，所以无法接收这个函数的返回值
	//下节讲如果进行通讯
	go func(a int, b int) bool {
		fmt.Println("a=", a, "b=", b)
		return true
	}(10, 20)

	//主程序循环等待
	for {
		time.Sleep(time.Second)
	}
}
