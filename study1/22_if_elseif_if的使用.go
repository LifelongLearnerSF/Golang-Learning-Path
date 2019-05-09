 package main
import "fmt"
func main(){
	// a := 10
	// if a == 11 {
	// 	fmt.Println("a == 10")
	// 	}else {	
	// 	fmt.Println("a != 10 ")		
	// }

	var a int
	fmt.Printf("请输入a的值:")
	fmt.Scanf("%d",&a)
	if a == 10 {
		fmt.Println("a == 10")
	}else if a > 10 {	
		fmt.Println("a > 10")		
	}else {
		fmt.Println("a < 10")
	}
	
}
