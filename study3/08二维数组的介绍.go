package main
import "fmt"
func main(){
	//有几个[]就是多少维,又有几次循环,例如下面[3][4]代表3行4列
	k:=0
	var a [3][4]int 
	
	for i:=0;i<3;i++{	
		for j:=0;j<4;j++{
				k++
				a[i][j]=k
				fmt.Printf("a[%d][%d]=%d",i,j,k)
			}
			fmt.Printf("\n")
		}-

	

	
	
}



