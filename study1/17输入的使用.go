package main
import "fmt"
func main(){

var a int
fmt.Printf("请输入变量a的值：")
fmt.Scan(&a)//阻塞用户输出，另外一种写法fmt.Scanf("%d",&a)
fmt.Printf("a=%d\n",a)
fmt.Printf("a=%c\n",a)



}
