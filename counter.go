package main
import"fmt"
func main(){
	var a,b,d float32
	var c string
	fmt.Println("这是一个简易的包含加减乘除的计算器")
	fmt.Println("请输入a")
	fmt.Scanf("%f\n",&a)
	fmt.Println("请输入符号")
	fmt.Scan(&c)
	fmt.Println("请输入b")
	fmt.Scanf("%f\n",&b)
	switch c{
	case "+":
		d=a+b
	fmt.Println("a加b的值为",d)
		case"-":
			d=a-b
			fmt.Println("a减b的值为",d)
    case"*":
    	d=a*b
    	fmt.Println("a乘b的值为",d)
    	case"/":
    		d=a/b
    		fmt.Println("a除以b的值为",d)
	default:
		fmt.Println("非法运算符")



	}




}


