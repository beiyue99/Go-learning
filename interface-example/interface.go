package main

import "fmt"

//go的interface类似void*，是个万能类型。interface{}表示空接口

type Book struct {
	name string
}

func Print(arg interface{}) {
	fmt.Println(arg)
	fmt.Println("===================")
	value, ok := arg.(Book)
	if !ok {
		fmt.Println("type not is Book\n")
	} else {
		fmt.Println("type is Book!\n")
	}
	fmt.Println("value is", value)
}

func main() {
	book := Book{"MyWorld"}
	Print(book) //不止结构体，interface可以接受任何类型参数

}

//注意： var a string="abcd"  底层是 pair<statictype:string,value:"abcd">
//var allType interface{}    allType=a    allType.(string)返回的第一个值就是通过pair的映射关系找到的，第二个为bool值
