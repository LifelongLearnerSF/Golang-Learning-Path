package main
import "fmt"
func main(){
var p *int
	p = new(int)
	*p=333
	fmt.Printf("*p=%d\n",*p)

	 q := new(int)
	 *q=666
	 fmt.Printf("*q=%d\n",*q)
}