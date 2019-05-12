package main

import "fmt"

//p是指向数组a,它是指向数组,是数组指针
//*p代表实参指向的内存,就是实参a
func modify(p *[5]int) {
	(*p)[0] = 7
	fmt.Printf("modify *p =%d\n", *p)
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}//初始化
	modify(&a)//地址传递
	fmt.Printf("main a =%d\n", a)
}
