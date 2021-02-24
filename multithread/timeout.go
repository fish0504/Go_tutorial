package main

import (
	"fmt"
	"time"
)
func fibonacci(c,quit chan int){
	x,y:=1,1

	for{
		select{
		case c<-x:
			x,y=y,x+y
		case <-quit:
			//cがブロックされたときに実行される
			fmt.Println("quit")
			return
		}
	}
}
func main(){
	c:=make(chan int)
	o:=make(chan bool)
	//quit:=make(chan int)
	//fibonacci(c,quit)
	go func(){
		for{
			select{
				case v:=<-c:
					fmt.Println(v)
				//タイムアウトを設定して,プログラム全体がブロックされるのを防ぐ
				case <-time.After(5*time.Second):
					fmt.Println("timeout")
					o<-true
					break
			}
		}
	}()
	<-o
}