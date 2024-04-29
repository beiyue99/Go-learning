package main

import (
	"fmt"
)

func main() {
	//声明一个变量，默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("type of a=%T\n", a)
	//给初始值
	var b int = 100
	fmt.Println("b=", b)
	fmt.Printf("type of b=%T\n", b)
	//自动匹配类型
	var c = 200
	fmt.Println("c=", c)
	fmt.Printf("type of c=%T\n", c)
	//方法四：最常用的方法，省去var
	//这种方法不支持声明全局变量
	e := 100
	fmt.Println("e=", e)
	fmt.Printf("type of e =%T\n", e)
	f := "abcd"
	fmt.Println("f=", f)
	fmt.Printf("type of f =%T\n", f)
	//声明多个变量
	var xx, yy int = 88, 99
	fmt.Println("xx=", xx, ",yy=", yy)
	var kk, ll = 33, "afwa"
	fmt.Println("kk=", kk, ",ll=", ll)
	//多行变量声明
	var (
		vv int  = 188
		jj bool = true
	)
	fmt.Println("vv=", vv, ",jj=", jj)

}
