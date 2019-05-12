package main

import "fmt"

//数组做参数,是值传递,是实参那每个元素拷贝一份给形参
func modify(a [5]int) {
	a[0] = 7
	fmt.Printf("modify a =%d\n", a)
}

//实参只是拷贝了一份给形参,原来的数值并没有变
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Printf("main a =%d\n", a)
}
