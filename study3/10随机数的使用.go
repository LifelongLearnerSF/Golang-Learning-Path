package main

import "fmt"
import "math/rand"
import "time"

func main() {
	//设置种子,如果种子参数一样,那么
	rand.Seed(time.Now().UnixNano())

	//产生随机数
	for i := 0; i < 6; i++ {
		fmt.Println("rand=", rand.Int()) //产生很大的数

	}
	fmt.Println("rand=", rand.Intn(55)) //限制在55以内的数字
}
