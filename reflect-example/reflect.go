package main

import (
	"fmt"
	"reflect"
)

//reflect提供了  reflect.TypeOf()  和reflect.ValueOf()获取值和类型

type User struct {
	Id   int
	Name string
	Age  int
}

func (this User) Call() {
	fmt.Println("user is Called...")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is ", inputType)        //输出main.User
	fmt.Println("inputType is ", inputType.Name()) //输出User
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("inputValue is ", inputValue) //输出 {1，"Aceld",19}
	fmt.Println("=============================")
	//通过inputType得到Numfield,遍历得到数据
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}

	//通过inputType获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s:%v\n", m.Name, m.Type)
	}

}

func main() {
	usr := User{1, "Aceld", 19}
	DoFiledAndMethod(usr)
}
