package main

//结构体的函数

import (
	"fmt"
)

// 类名或者成员属性名大小写作为共有和私有的依据
type Animal struct {
	Name string
	age  int
}

func (this *Animal) show() {
	fmt.Println("Name=", this.Name)
	fmt.Println("Age=", this.age)
}

// 用指针传参，否则默认值传递
func (this *Animal) setName(name string) {
	this.Name = name
}

func main() {
	mycat := Animal{"Tom", 22}
	mycat.show()
	mycat.setName("Bulu")
	mycat.show()
	//因为在同一个包，即使age为小写，也可以访问的到
	fmt.Println("mycat Name=", mycat.Name, "age=", mycat.age)
}
