package main

import "fmt"

func main() {
	//声明map
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 is null")
	}
	//同理，没有空间不可使用，需要开辟空间
	map1 = make(map[string]string, 10)
	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "python"
	fmt.Println(map1)
	//第二中声明方式,make就是创建map，不用指定大小，也可以直接使用
	map2 := make(map[int]string)
	map2[1] = "apple"
	map2[2] = "bnana"
	map2[3] = "peer"
	fmt.Println(map2)
	//第三种声明方式,也是创建了map，可以直接用
	map3 := map[int]string{
		1: "Shanghai",
		2: "Beijing",
		3: "Wuhan",
	}
	map3[4] = "Shenzhen"
	fmt.Println(map3)
}
