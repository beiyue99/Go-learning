package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Wirter interface {
	WirterBook()
}

type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Raed a Book!")
}

func (this *Book) WirterBook() {
	fmt.Println("Wirte a Book!")
}

func main() {
	b := &Book{}
	//b  pair<type:Book*,value:Book的地址>
	var r Reader
	fmt.Printf("type of r is%T\n", r) //输出nil
	fmt.Printf("r is %v\n", r)        //输出nil
	r = b
	fmt.Printf("type of r is%T\n", r) //输出*Book
	fmt.Printf("r is %v\n", r)        //输出  &{},即某对象的地址，在这里是Book对象
	//r  pair<type:Book*,value:Book的地址>
	r.ReadBook()
	fmt.Println("+++++++++++++++++++++++++")
	value, ok := r.(Wirter)
	//其实也好理解，r被赋值后是&{}类型，也就是对象指针，而Wirter也是对象指针类型，所以这里断言成功,返回的值就是对象指针
	fmt.Println(ok)                //返回true
	fmt.Println("value is", value) //输出&{}
	var w Wirter = value           //返回的这个&{}对象指针可以给w赋值
	w.WirterBook()
}
