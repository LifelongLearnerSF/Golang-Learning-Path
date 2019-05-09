package main
import(
	"fmt"

)

//全局变量在函数外部使用(注意全局变量不能用:=声明赋值)
var a int

func main() {
		a:=1//再次声明赋值
		{ 
	
		a :=2//注意:  a=2和a :=2 输出结果是不一样的
			
	       defer fmt.Println(a)


	}
		
	
	fmt.Println(a)
}
