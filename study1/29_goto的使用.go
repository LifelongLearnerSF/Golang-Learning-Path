package main
import "fmt"
func main(){
	fmt.Println("11")
	goto END
	fmt.Println("22")

END:
	fmt.Println("33")
}