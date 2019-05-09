package main

import  "fmt"


func main() {
    fmt.Println("1-100之间的质数为:")
    // i应直接从2开始
    for i := 2; i <= 100; i++ {
        for n := 2; n <=i; n++ {
            // 当走到最后n等于i 了，则说明下面的i%n==0 && n < i 始终没有成立。说是这个数是个质数。
            if n == i {
                fmt.Printf("%d ", i)
            }
            // 当满足这个条件的时候就终止里面的循环，不用继续往下走了,因为它已经不是一个质数了。
            if i%n == 0 && n < i {
                break
            }
        }
    }
}