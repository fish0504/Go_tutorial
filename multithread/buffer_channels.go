package main

import "fmt"


/*
ch:=make(chan type,value)
value==0 ! //バッファリング無し
value >0 ! //バッファリング
*/
func main(){
	c:=make(chan int,1)
	c<-1
	c<-2
	fmt.Println(<-c)
	fmt.Println(<-c)
}