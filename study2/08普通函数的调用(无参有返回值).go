package main
import "fmt"

func funca()(a int){
	
	a=111
	fmt.Println("funca a=",a)
	b:=funcb()
	fmt.Println("funcab b=",b)
	return
}


func funcb()(b int){
	b=222
	fmt.Println("funcb b=",b)
	return
}


func main(){
	a :=funca()
  fmt.Println("main a =",a)
}





