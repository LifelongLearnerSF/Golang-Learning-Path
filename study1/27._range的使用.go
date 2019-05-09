package main
import "fmt"
func main(){
	str := "abcd"
	for i := 0;i < len(str);i++{
		fmt.Printf("str[%d]=%c\n",i,str[i])
	}
	
	str2 := "efg"
	for date , value := range str2{//range的作用:索引位置,位置上的"值"
		fmt.Printf("str2[%d]=%c\n",date ,value)
	} 



	}




	





