package main

import "fmt"

type Human struct {
	name string
	sex  string
}

// 定义超人类，继承自Human
type Superman struct {
	Human
	lelve int
}

func (this *Human) Eat() {
	fmt.Println("Human Eat!")
}

func (this *Human) Walk() {
	fmt.Println("Human Walk!")
}

func (this *Superman) Fly() {
	fmt.Println("Superman Fly!")
}

func (this *Superman) Eat() {
	fmt.Println("Superman Eat cake !")
}

func main() {
	h1 := Human{"zhangsan", "female"}
	h1.Walk()
	h1.Eat()
	Superh1 := Superman{Human{"dagu", "male"}, 999}
	fmt.Println("++++++++++++++++++++++++")
	Superh1.Walk()
	Superh1.Fly()
	Superh1.Eat()
	//也可以先var s Superman  然后通过.号赋值三个属性
}
