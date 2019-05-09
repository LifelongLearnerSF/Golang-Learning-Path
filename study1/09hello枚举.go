package main
import "fmt"
func main(){
    const(
		a = iota
        
		b=iota

	)
	fmt.Printf("a=%d,b=%d\n",a,b)
	const(
		c = iota
		d = iota
		e = iota

	)
	fmt.Printf("a=%d,b=%d,e=%d\n",a,b,e)
	const i  =iota
	fmt.Println("i=",i)











}