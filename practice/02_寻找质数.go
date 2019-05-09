package main
import "fmt"
func main(){
 //判断一个数是否为质数,只要不被2和本身之间的数整除即可
    fmt.Println("1-100之间的质数为:")
for i:=2;i <=100;i++{
    // fmt.Println("1-100之间的质数为:")
    for n:= 2;n<=i;n++{
            if n == i{
                fmt.Printf("%d",i)
            }
            if i%n == 0 && n<i{
         
                 break
            }
    




        }
    }

}