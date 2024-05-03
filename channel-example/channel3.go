package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		//如果不关闭，没数据还读就会出现问题

	}()

	//for {
	//if data, ok := <-c; ok {
	//fmt.Println(data)
	//} else {
	//break
	//}
	//}
	fmt.Println("===========================")
	//下面的循环跟上面的循环作用一样
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Main finished!")

}
