package main
import(
	"fmt"

)

//在不同的作用域变量名可以同名
//就近原则
var a byte

func main() {
	var a int
	a=1
	fmt.Printf("1=%d\n",a)
	{a :=2
	fmt.Printf("2a=%d\n",a)}

	fmt.Printf("3=%d\n",a)
	Test(a)
}
func Test(a int){
	fmt.Println(a)
}


