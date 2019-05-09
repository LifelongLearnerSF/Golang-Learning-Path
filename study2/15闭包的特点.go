package main
import "fmt"
func main(){
	
	a:=myfunc()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}

func myfunc() func()int {
	x:=1
	return func()int{
	x++
	return x*x
}



}






