package main
import "fmt"


func fibonacci() func(int) int{
	f1:=0
	f2:=1
	return func(x int)int{
		sum:=f1+f2
		f1=f2
		f2=sum
		return sum
	}
}


func main(){

	f:=fibonacci()

	for i:=0;i<10;i++{
		fmt.Println(f(i))
	}
}