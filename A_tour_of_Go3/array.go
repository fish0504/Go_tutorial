package main
import "fmt"


func main(){
	var a [2]string //string a[2](c++)
	b:=[2]string{"hello","micky"} //omittion var definition
	fmt.Println(b[0],b[1])
	a[0]="Hello"
	a[1]="world"
	

	fmt.Println(a)

	primes:=[6]int{1,3,5,7,11,13}
	fmt.Println(primes)
}