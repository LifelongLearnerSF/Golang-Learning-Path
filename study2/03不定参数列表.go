package main
import "fmt"
func main(){
	MyFunc2(1,2,3)

}

func MyFunc2(a ...int){
for date,value := range a{

	fmt.Printf("a[%d]=%d\n",date,value)
}





}

