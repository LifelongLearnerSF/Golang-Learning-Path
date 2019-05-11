package main
import "fmt"
func main(){
	//数组之间的比较条件:同类型,长度一样,只能==或者!=
	a :=[5]int{1,2,3,4,5}
	b :=[5]int{1,2,3,4,5}
	c :=[5]int{1,2,3,}
	/*数组比较*/
	fmt.Println("a==b",a==b)
	fmt.Println("a==c",a==c)
	fmt.Println("a!=c",a!=c)	

	//数组赋值,把a的值赋给d
	var d [5]int
	d=a
	fmt.Println("d=",d)
	
}



