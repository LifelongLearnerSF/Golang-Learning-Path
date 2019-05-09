package main
import "fmt"
func main() {
	//defer的作用是延时调用,在main函数结束前一瞬间调用
	defer fmt.Println("b")
	fmt.Println("a")
}
