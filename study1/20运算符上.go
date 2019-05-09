 package main
import "fmt"
func main(){
	 a:=3
	 b:=3


	fmt.Println("a>b结果是:",a>b)
	//关系运算符>  <  =>

fmt.Println("4!=3的结果是:",!(4!=3))
//!逻辑运算符 非

fmt.Println("true && true 的结果是：",true && true)
fmt.Println("true && false 的结果是：",true && false)
//&&逻辑运算符 与

fmt.Println("true || false 的结果是：",true || false)
fmt.Println("true || true的结果是：",true || true)
fmt.Println("false || false的结果是：",false || false)
 //||逻辑运算符 或 只有都为假是才是假

fmt.Println("4>3 || 4==3 的结果是：",4>3 || 4==3 )
fmt.Println("4>3 && 4==3 的结果是：",4>3 && 4==3 )

a=10
fmt.Println("8<a &&a<11 的结果为：",!(a>8 && a<11))
//不可以写成8<a<11的形式
}
