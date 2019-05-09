package main
import "fmt"
func main(){
ch:='a'
str:="a"
//字符串是由多个字符组成，单引号和双引号的区别
fmt.Println("ch=", ch)
fmt.Printf("ch以字符方式打印:%d\n", ch)
//上面两种打印效果是一样的
fmt.Println("str=", str)

str2:="shufang"
fmt.Printf("str的类型是: %T\n",str2)
fmt.Println("len(str2)=",len(str2))
fmt.Printf("str2[0]=%c str2[5]=%c\n",str2[0],str2[5])
fmt.Println(str2)

}