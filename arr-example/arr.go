package main

import "fmt"

func main() {
	var myArr1 [10]int
	myArr2 := [10]int{1, 2, 3, 4}
	for i := 0; i < len(myArr1); i++ {
		fmt.Println(myArr1[i])
	}
	for index, value := range myArr2 {
		fmt.Println("index=", index, "value=", value)
	}

}
