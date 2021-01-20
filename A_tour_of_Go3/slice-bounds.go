package main

import "fmt"

func main(){
	s:=[]int{1,2,3,4,1,3}

	s =s[1:4]
	fmt.Println(s)//[2,3,4]
	s=s[:2]			
	fmt.Println(s)//[2,3]
	s=s[1:]

	fmt.Println(s)//[3]
	//cap(s) 配列の長さを返す
	fmt.Println(cap(s))
}