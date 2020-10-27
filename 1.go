package main
import "fmt"
func main() {
	a := []int{12, 23, -43, 3, 67, 10,-2,1}
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			if a[j]<0{
				a[j]=-a[j]
				}
				}
			}
		}
		fmt.Print(a)
		}





