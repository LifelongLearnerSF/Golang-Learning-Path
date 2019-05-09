package main
import "fmt"
func main(){
	var score int

	fmt.Printf("请输入分数：")
	fmt.Scanf("%d",&score)
	switch {
	case score <60:
		fmt.Println("不及格")
		
	case  score >=60&&score<=80:
		fmt.Println("良好")
		
	default:
		fmt.Println("优秀")
	
	
	



	}




	




}
