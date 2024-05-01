package main

import "fmt"

func main() {
	// 切片的追加操作
	var arr []int
	//arr[0] = 1,会出错，因为没有开辟空间
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) //0,0
	arr = append(arr, 1)
	fmt.Printf("len=%d,cap=%d\n", len(arr), cap(arr)) //1,1
	//跟vector一样，会扩容
	//切片的截取操作
	s := []int{1, 2, 3} //len=3,cap=3

	s1 := s[0:2] //左闭右开，取出第0和第一个元素  1，2
	fmt.Println(s1)
	//要注意的是s1和s指向相同的内存块
}
