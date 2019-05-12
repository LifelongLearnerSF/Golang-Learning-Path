package main

import "fmt"

func main() {
	//数组[]是固定的,即长度和容量是固定的不可以更改
	a := [5]int{}
	fmt.Println("a=", a)
	fmt.Printf("len(a)=%d,cap(a)=%d\n", len(a), cap(a))
	//切片[]是空的,len和cap是可以改变的
	b := []int{}
	fmt.Println("b=", b)
	fmt.Printf("len(b)=%d,cap(b)=%d\n", len(b), cap(b))

	b = append(b, 123) //给切片末尾追加一个成员
	fmt.Println("b=", b)
	fmt.Printf("append.len(b)=%d,cap(b)=%d\n", len(b), cap(b))
}
