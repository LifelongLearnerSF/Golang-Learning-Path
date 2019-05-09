package main
import "fmt"
func funcc(c int){
	fmt.Println("c=",c)

}



func funcb(b int){
	funcc(b-1)
	fmt.Println("b=",b)
}
func funca(a int){
	funcb(a-1)
	fmt.Println("a=",a)


}

func main(){
	funca(8)
	fmt.Println("main")
	
}





