package main 

import "fmt"


// var p *int

//i:=42
//p=&i

//fmt.Println(*p)
//*p=21


func main(){
	i:=42

	p:=&i
	
	fmt.Println(*p)
	*p=21
	fmt.Println(i)
	fmt.Println(2701/37)
}