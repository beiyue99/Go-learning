package main

import(
	"Golang-learn/init/lib1"

	"Golang-learn/init/lib2"
)
//在 Go 语言中，为了从其他包中访问一个函数或变量，它必须是以大写字母开头的，这样它才是公开可导出的。
func main(){
	lib1.Lib1Test()
	lib2.Lib2Test()
}










