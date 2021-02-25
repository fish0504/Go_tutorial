package main

import "fmt"

func main(){


	var x int32
	fmt.Scan(&x)

	x=x%100
	fmt.Println(100-x)
}