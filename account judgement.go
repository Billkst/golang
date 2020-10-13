package main

import (
	"fmt"
	"os"
)
func main() {
	fmt.Println("欢迎来到账号登陆系统")
	fmt.Println("请输入账号")//账号：ljx 密码：asd123
	fileName := "D:/account/a.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer file.Close()
	id:= make([]byte,3,3)
	code:=make([]byte,6,6)
	n, err := file.Read(id)
	n, err = file.Read(code)
	n = 1
	var a,b string
	fmt.Scanln(&a)
	fmt.Println("请输入密码")
	fmt.Scanln(&b)
	if b==string(code) {
		fmt.Println("密码正确")
	} else{
		fmt.Println("密码错误")
	}

	fmt.Println(n)

}