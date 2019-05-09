package main
import "fmt"
func main(){
	a :=666
	str := "abc"
	func(){
	
		//闭包以引用的方式捕获外部变量
		fmt.Printf("内部:a=%d,str=%s\n",a,str)
	}()//()代表直接引用
	

	fmt.Printf("外部:a=%d,str=%s\n",a,str)
	
}


