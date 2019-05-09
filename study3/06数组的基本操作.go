//打印班级50个学生id
package main
import "fmt"
func main(){
	//声明,元素的个数必须是常量
	var a [10] int
	//赋值
	for i:=0;i<len(a);i++{
		a[i] = i+1  
		
	 }
	 //打印
	 for date,value := range a{
		 fmt.Printf("a[%d]=%d\n",date,value)
	 }



}



