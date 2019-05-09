package main
import "fmt"
func main(){
var a int =10

fmt.Printf("&a=%p,a=%d\n",&a,a)
var p *int
p = &a
fmt.Printf("p=%p,&=%p\n",p,&a)
*p=22

fmt.Printf("a=%d,&=%p\n",a,&a)
}