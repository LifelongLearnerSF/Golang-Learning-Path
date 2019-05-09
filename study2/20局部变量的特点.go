package main
import(
	"fmt"

)

//局部变量只能在函数内部使用
//作用域:变量作用的范围
func main() {
	var a int
	a =2
		{ 
		  a =1
		 
	      defer  fmt.Println(a)



	}
	
	fmt.Println(a)
}
