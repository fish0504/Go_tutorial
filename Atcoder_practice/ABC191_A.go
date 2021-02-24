package main

import "fmt"


func main(){
	var v,t,s,d float64
	fmt.Scan(&v,&t,&s,&d)

	start:=t*v
	end:=s*v
	if start<=d && d<=end{
		fmt.Println("No")
	}else{
		fmt.Println("Yes")
	}


	
}