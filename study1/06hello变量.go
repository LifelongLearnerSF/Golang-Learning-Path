package main 
import "fmt"
func test()(a,b,c int){
	return 1,2,3
}
	
func main(){
	a,b:=10,20
	
	a,b = b,a
	fmt.Printf("a=%d,b=%d\n",a,b)
	i,j:=10,20
	var tmp int
	tmp, _=i, j
	fmt.Println("tmp =",tmp)

	var c,d,e int
	c,d,e=test()
	fmt.Printf("c=%d,c%d,e=%d\n",c,d,e)


















}


