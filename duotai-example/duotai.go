package main

//先说一下原理
//c++中，多台实现方式是通过父类指针指向子类对象实现多态
//go语言中，是通过interface抽象类，子类重写抽象类方法，
//然后这个interface指针指向子类实现多态

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}

func Print(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}

func main() {
	var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep()
	animal = &Dog{"black"}
	animal.Sleep()
	fmt.Println("===============")
	Print(&Cat{"red"})
	Print(&Dog{"blue"})
}
