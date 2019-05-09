package main
import "fmt"


func funca(a int){
	if a ==1 {
		fmt.Println("a=",a)
		return
	}
	funca(a-1)
	fmt.Println(a)

}



func main(){
	funca(4)
}





