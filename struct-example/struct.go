package main

import "fmt"

// type跟c++的typedef区别不大，但是定义相反  type myint int
// 结构体的基础用法
type Book struct {
	title string
	auth  string
}

func Modify(mybook *Book) {
	mybook.auth = "lisi"
}

// 结构体传参传递的是指而不是引用
func main() {

	var Mybook Book
	Mybook.title = "Golang"
	Mybook.auth = "zhangsan"
	fmt.Println(Mybook)
	Modify(&Mybook)
	fmt.Println(Mybook)
}
