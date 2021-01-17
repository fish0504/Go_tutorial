package main

import (
	"fmt"
)

func Sqrt(x float64)(z float64){
	z=1.0
	time:=int64(0)
	for i:=1;i<20;i++{
		//now:=float64()
		// if ((z*z-x)/(2*z))<0.01{
		// 	time=int64(i)
		// 	break
		// }
		z-=(z*z-x)/(2*z)
	}
	fmt.Println(time)
	return z
}

func main(){
	fmt.Println(Sqrt(16))
}