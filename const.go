
package main

import(
	"fmt"
)

//const 可以定义枚举类型
const (
	//添加一个关键字iota，初始值是0，逐渐累加，不用自己加了
	BeiJing=iota
	ShangHai
	ShenZhen
	GuangZhou
)





func main(){
	//常量
	const length int =10
	fmt.Println("length=",length)
	//length=100   //常量不允许修改，此处会报错
	fmt.Println("BeiJing=",BeiJing)
	fmt.Println("ShangHai=",ShangHai)
}




