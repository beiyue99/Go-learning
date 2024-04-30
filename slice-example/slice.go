package main

import "fmt"

// 切片的定义方式,注意切片传参的时候是作为引用传递
func main() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("len=%d,slice1=%v\n", len(slice1), slice1)
	//len=1,2,3  slice1=[1,2,3]
	var slice2 []int //此时slice2没有空间，如果直接赋值会出错
	slice2 = make([]int, 3)
	fmt.Printf("len=%d,slice2=%v\n", len(slice2), slice2)
	slice2[0] = 99
	fmt.Printf("len=%d,slice2=%v\n", len(slice2), slice2)
	//var slice3 = make([]int, 3)
	//slice4 := make([]int, 3)
	//判空： if(slice4==nil)
}
