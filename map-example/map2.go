package main

import "fmt"

func Modify(citymap map[string]string) {
	citymap["England"] = "Baigong"
}

func main() {
	citymap := make(map[string]string)
	citymap["china"] = "Beijing"
	citymap["japan"] = "Tokyo"
	citymap["USA"] = "NewYork"
	//遍历
	for key, value := range citymap {
		fmt.Println("key=", key, "value=", value)
	}
	//删除
	delete(citymap, "china")
	//修改
	citymap["USA"] = "Huashengdun"
	fmt.Println("==========")
	//遍历
	for key, value := range citymap {
		fmt.Println("key=", key, "value=", value)
	}
	//要注意，跟切片一样，函数传参传递的都是引用
	fmt.Println("==========")
	Modify(citymap)
	//遍历
	for key, value := range citymap {
		fmt.Println("key=", key, "value=", value)
	}
}
