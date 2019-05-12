package main

import "fmt"

func main() {
	//自动推导类型,同时初始化
a1 := []int{1,2,3,4,5}
fmt.Println("a1=",a1)
//借助make函数,make函数格式(数组类型,长度,容量)
a2 := make([]int,3,5)
fmt.Printf("len=%d,cap=%d\n",len(a2),cap(a2))
//make函数没有指明容量的时候,容量和长度大小一样
a3 := make([]int,3)
fmt.Printf("len=%d,cap=%d\n",len(a3),cap(a3))
}
