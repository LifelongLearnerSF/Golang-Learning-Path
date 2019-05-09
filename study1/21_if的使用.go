 package main
import "fmt"
func main(){
	var s int
	fmt.Printf("请输入变量a的值：")
	fmt.Scanf("%d",&s)

	if s==10{
		fmt.Println("正确")


	}


	//如果为真，则输出a == 1
	if s!=10{
		fmt.Println("错误")
	}

}
