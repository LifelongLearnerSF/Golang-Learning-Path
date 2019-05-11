package main
import "fmt"
func main(){
	//初始化,声明的同时赋值
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a=",a)
	
	//元素个数不能大于数组长度
	var b [5]int = [5]int{1,2,3}
	fmt.Println("b=",b)

	var c [5]int = [5]int{1:10,3:100}
	fmt.Println("c=",c)

}



