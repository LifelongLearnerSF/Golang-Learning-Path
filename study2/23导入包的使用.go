package main
// //1传统导入
// import "fmt"
// import "os"
// func main(){
// 	fmt.Println("fmt")
// 	fmt.Println("os=",os.Args)
// }




// //2()导入多个包
// import (
// 	"fmt"
// 	"os"
// )
// func main(){
// 	fmt.Println("fmt")
// 	fmt.Println("os=",os.Args)
// }


// //3  .操作导入包,就不需要再次写包名,但是这种用法不建议使用
// import (
// 	."fmt"
// 	."os"
// )
// func main(){
// 	Println("fmt")
// 	Println("os=",Args)
// }


// //4别名操作导入包
// import (
// 	a "fmt"
// 	b "os"
// )
// func main(){
// 	a.Println("fmt")
// 	a.Println("os=",b.Args)
// }


//5  _操作导入包
import (
	_ "fmt"
	_ "os"
)
func main(){
	
}
