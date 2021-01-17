package main
import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool =false	
	MaxInt uint64 = 1<<64-1
	Maxint64 int64=1<<32
	z complex128=cmplx.Sqrt(-5+12i)
	floatMax float32=3.1415926535
	str string ="Take it role!"
)

func main(){
	fmt.Printf("Type: %T Value: %v\n",ToBe,ToBe)
	fmt.Printf("Type: %T Value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type: %T Value: %v\n",z,z)
	fmt.Printf("Type: %T Value: %v\n",Maxint64,Maxint64)
	fmt.Printf("Type: %T Value: %v\n",floatMax,floatMax)
	fmt.Printf("Type: %T Value: %v\n",str,str)

}