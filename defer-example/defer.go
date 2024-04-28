package main

import"fmt"

func main(){

//类似c++的析构函数，程序结束执行,如果下面还有一个defer，按照先进后出（栈）执行
	defer fmt.Println("main end!")
	fmt.Println("main::hello1")
	fmt.Println("main::hello1")
}



