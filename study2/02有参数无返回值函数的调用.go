package main
import "fmt"

func MyFunc1(a int,b bool,c float32){//形参
	fmt.Printf("a=%d,b=%t,c=%f\n",a,b,c)
}


func main(){
 MyFunc1(100,true,3.14)//实参

}


