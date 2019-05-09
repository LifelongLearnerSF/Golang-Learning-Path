package main
import "fmt"

func Add(a,b int)int{
return a + b

}

//My是一个函数类型
type My func(int,int)int//没有函数名,没有{}

func main(){
	var You My
	You = Add
	result:=You(10,20)
	fmt.Println("result=",result)
	

}



