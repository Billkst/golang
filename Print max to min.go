package main//随机生成五个整数并由大到小排序

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	fmt.Println("随机生成5个正数(100以内）并由大到小排序")
	var a [5]int
	n:=len(a)
	for i:=0;i<n;i++ {
	a[i]=rand.Intn(100)
	}
	for i:=0;i<n-1;i++{
	for j:=0;j<n-1-i;j++{
	if a[j]<a[j+1]{
		a[j],a[j+1]=a[j+1],a[j]
	}
	}
	}
	for i:=0;i<n;i++ {
		fmt.Printf("%d\n", a[i])
	}
}
