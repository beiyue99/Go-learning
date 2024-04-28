package main
import(
	_"Golang-learn/init-example/lib1"
	"Golang-learn/init-example/lib2"
//mylib2 "...lib2"  起别名
//. "...lib2" 导入当前文件，不需要通过包名调用包内方法
)
//在 Go 语言中，为了从其他包中访问一个函数或变量，它必须是以大写字母开头的，这样它才是公开可导出的。
func main(){
	//lib1.Lib1Test()
	//如果不使用lib1的函数，但是导入lib1，会编译错误，解决方案是加一个_
	lib2.Lib2Test()
}










