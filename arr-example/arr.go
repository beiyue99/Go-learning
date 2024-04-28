package main


import "fmt"

func main(){
	var arr1[10]int

	arr2:=[10]int{1,2,3,4}


	//for循环内部不可以加（）
	for i:=0;i<len(arr1);i++{
		fmt.Println(arr1[i])
	}
	for index,value:=range arr2{
		fmt.Println("index=",i


}
