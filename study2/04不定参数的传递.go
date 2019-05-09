// package main
// import "fmt"
// func main(){
// 	MyFunc1(`b`,`c`,`d`)

// }

// func MyFunc1(a ...string){
// for date,value := range a{

// 	fmt.Printf("a[%d]=%s\n",date,value)
// }





// }
package main
import "fmt"
func main(){
	MyFunc1("bcd")

}

func MyFunc1(a ...string){
for date,value := range a{

	fmt.Printf("a[%d]=%s\n",date,value)
}





}