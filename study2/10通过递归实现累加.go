package main
import "fmt"
func main(){
 num :=test(100)
 fmt.Printf("num=%d",num)
}
func test(i int)(sum int){
	
	if i ==1{
		return 1
	}
	return  i+test(i-1)
}





