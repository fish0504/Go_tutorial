package main

import (
	"fmt"
	"math"
)

func pow(x,n,lim float64)float64{
	//if文の条件の前に宣言できる
	if v:=math.Pow(x,n);v<lim{
		return v
	} else{
		fmt.Printf("%g >= %g\n",v,lim)
		return lim
	}
	//can't use v here though
	
}

func main(){
	fmt.Println(
		pow(3,2,10),
		pow(3,3,20),
		pow(3,4,30),
	)
}