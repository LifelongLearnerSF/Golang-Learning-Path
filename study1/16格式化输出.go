package main
import "fmt"
func main(){
a:=10
b:="abc"
c:='A'
d:=3.14
fmt.Printf("a=%d,b=%s,c=%c,d=%f\n",a,b,c,d)
fmt.Printf("a=%v,b=%v,c=%v,d=%v\n",a,b,c,d)
fmt.Printf("a=%T,b=%T,c=%T,d=%T\n",a,b,c,d)
//%d整型格式
//%s字符串格式
//%c字符格式
//%f浮点型格式
//%Tgo语法的输出值的类型




}
