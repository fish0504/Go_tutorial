package main 

import (
	"fmt"
	"math"
)
// variable v ,Type T
// T(v):= convert v to Type T

func main(){
	var x,y int =4,5
	var f float64 =math.Sqrt(float64(x*x+y*y))

	var z uint=uint(f)
	fmt.Println(x,y,z,f)

	// var i int =42
	// var flo float32=float32(i)
	// var u uint=uint(flo)

	i:=42
	flo:=float64(i)
	u:=uint(flo)
	fmt.Printf("i:%8d flo:%.3f u :%v\n",i,flo,u)

}