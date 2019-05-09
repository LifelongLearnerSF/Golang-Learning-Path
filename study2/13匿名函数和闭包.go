package main
import "fmt"
func main(){
	a := 10
	str := "bcd"
	//匿名函数,没有函数名字,函数定义,
	F1:=func(){//自动推导类型
		fmt.Printf("a=%d,,str=%s\n",a,str)
	}
	F1()
	

	//给一个函数类型起名(注意不是给函数起名)
	type Test1 func()
	var F2 Test1
	F2 =F1
	F2()
                                   


	//定义匿名变量,同时声明
	func(){
		fmt.Printf("a=%d\n str=%s\n",a,str)
	}()

	//带参数的匿名函数
	func(i,j int){
		fmt.Printf("i=%d,j=%d\n",i,j)
	
	}(1,2)
	
	
	//带参带返回值
	x,y:=func(i,j int)(max,min int){
		if i>j{
			max =i
			min =j
		}else{
			max =j
			min =i
		}
		return 
	}(10,20)

	fmt.Printf("max=%d,min%d\n",x,y)
}



