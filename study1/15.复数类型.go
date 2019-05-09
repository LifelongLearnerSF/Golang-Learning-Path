package main
import "fmt"
func main(){
tc:= 3.14 + 5i
fmt.Println("tc=",tc)
fmt.Printf("tc type is %T\n",tc)
fmt.Println("real(tc)=",real(tc),"imag(tc)=",imag(tc))

}