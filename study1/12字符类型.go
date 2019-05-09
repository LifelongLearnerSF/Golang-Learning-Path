package main
import "fmt"
func main(){
   var tmp byte 
   tmp= 65
   fmt.Printf("tmp =  %c\n",tmp)
   fmt.Printf("tmp =  %d\n",tmp)
//格式化输出%c以字符方式打印，%d以整型方式打印



tmo :='a'
//字符类型以单引号''
fmt.Printf("以字符类型打印： %c\n",tmo)
fmt.Printf("以整型打印： %d\n",tmo)





fmt.Printf("大写:%c,小写:%c\n",'A','a')
//fmt.Printf("大写:%d,小写:%d\n",'A','a')
//fmt.Printf("大写转小写: %d, %c\n",'A'-32,'A'-32)
fmt.Printf("大写转小写:%c\n",'A'+32)
fmt.Printf("小写转大写:%c\n",'a'-32)







}