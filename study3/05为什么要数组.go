//打印班级50个学生id
package main
import "fmt"
func main(){
	var name [50] int
	for i:=0;i<len(name);i++{
		name[i]=i+1
		fmt.Printf("name[%d]=%d\n",i,name[i])
	}
	}



