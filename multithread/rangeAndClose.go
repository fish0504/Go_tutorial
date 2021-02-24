package main

import (
	"fmt"
)

func fibonacci(n int, c chan int){
	x,y:=1,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}

func main(){
	//10個のint型要素を持てるチャンネルを作成
	c:=make(chan int,10)
	go fibonacci(cap(c),c)
	v,ok:= <-c
	
	for i:=range c{
		fmt.Println(v,ok)
		v,ok=<-c
		fmt.Println(i)
	}
}