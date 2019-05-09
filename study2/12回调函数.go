package main
import "fmt"

func Add(a , b int)int{
	return a+b
}
type Test1 func(int,int)int
func Calo(a,b int,Test2 Test1)(result int){
	
	result=Test2(a,b)
	return


}


func main(){
	c :=Calo(10,20,Add)
	fmt.Println("c=",c)

	

}



