package main

import (
	"fmt"
)


func fool(a string,b int) int {
	fmt.Println("a=",a)
	return 22
}

func fool2(a string,b int)(int,int){
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	return 666,777
}

//也可以写成   r1,r2 int 表示他俩都是int
func fool3(a string,b int)(r1 int,r2 int){
	fmt.Println("a=",a)
	fmt.Println("b=",b)
	r1=1000
	r2=2000
	return
}


func main(){
	c:=fool("abc",555)
	fmt.Println("c=",c)
	ret1,ret2:=fool2("haha",999)
	fmt.Println("ret1=",ret1,"ret2=",ret2)
	ret1,ret2=fool3("foo3",333)
	fmt.Println("ret1=",ret1,"ret2=",ret2)
}









