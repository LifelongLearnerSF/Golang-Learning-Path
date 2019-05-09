 package main
import "fmt"
func main(){
	var num int
	fmt.Printf("请按下数字：")
	fmt.Scanf("%d",&num)
	switch num{
	case 1:
		fmt.Println("显示数字:",1)
		fallthrough
	case  2:
		fmt.Println("显示数字:",2)
		fallthrough
	case 3: 
		fmt.Println("显示数字:",3)
	    fallthrough
	default :
		fmt.Println("显示其它")
	
		//fallthrough 表示不跳出switch语句


	}




	




}
