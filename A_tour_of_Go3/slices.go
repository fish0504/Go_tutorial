package main
import "fmt"


//slices:=lenght variable array

//ex)	a[low:high]:= [low,high) lowは含むがhighは含まない
func main() {
	primes:=[6]int {1,1,2,3,5,8}

	var s []int=primes[1:4]// primesのprimes[1]~primes[3]
	fmt.Println(s)

}