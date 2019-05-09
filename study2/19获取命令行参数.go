package main
import(
	"fmt"
	"os"
)


func main() {
	//接收用户传递的参数,都是以字符串方式传递
	for date,value := range os.Args{
		fmt.Printf("参数[%d]:%s\n",date,value)
	}

}
