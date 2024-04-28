package main

import "fmt"


func swap1(a* int ,b* int){
	var temp int
	temp=*a
	*a=*b
	*b=temp
}



func main(){
	var a int =10
	var b int =20
	swap1(&a,&b)
	fmt.Println("a=",a,"b=",b)
}
