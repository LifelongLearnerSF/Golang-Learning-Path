package main
import "fmt"



func main() {
	 a := 111
	 b := 222
	 //无参数无返回值
	defer func(){
		fmt.Printf("a=%d,b=%d\n",a,b)
		
	}()
	//有参数
	defer func(a,b int){
		fmt.Printf("a=%d,b=%d\n",a,b)
		
	}(a,b)
		a = 333
	 	b = 444
		fmt.Printf("a=%d\n,b=%d\n",a,b)
}
