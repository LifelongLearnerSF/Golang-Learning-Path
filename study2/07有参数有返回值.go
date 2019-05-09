package main
import "fmt"

func maxandmin(a ,b int)(max,min int){
	if a>b{
		max = a
		min =b
	}else{
		min = a
		max = b
	}
	return


}


func main(){


	x,y:=maxandmin(10,20)

	fmt.Printf("max=%d,min=%d\n",x,y)
	 _,c:=maxandmin(10,20)
	 fmt.Printf("c=%d\n",c)


}






