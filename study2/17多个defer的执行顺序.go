package main
import "fmt"

func test(x int){
	
	rsult :=100/x
	fmt.Println(rsult)
}

//多个defer使用的时候,"先进后出",遇到不可执行的函数或者延时调用,这些调用依旧会被执行
func main() {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer test(0)
	defer fmt.Println("c")
}
