package main

import "fmt"

/*
select:= channel上のデータを監視する




*/
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
	//int型のchannelを作成
	c:=make(chan int)

	quit:=make(chan int)

	go func(){
		for i:=0;i<10;i++{
			fmt.Println(<-c)
		}
		quit<-0
	}()
	fibonacci(c,quit)
}
/*
select{
	case i:=<-c:
	//use i
	default:
	//cがブロックされたときここが実行される
}
*/