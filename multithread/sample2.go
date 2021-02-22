package main

import "fmt"

/*
共有されたメモリへのアクセスは必ず同期されていなければならない
channelを通して通信を行う
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})

*/

func sum(a []int ,c chan int){
	total := 0
	for _,v:=range a{
		total+=v
	}
	c<-total	//send total to c
}
/*
channekへの送信と受け取り方
ch <- v    // vをchannel chに送る。
v := <-ch  // chの中からデータを受け取り、vに代入する。
*/

func main(){
	a :=[]int {7,2,8,-9,4,0}

	//cというint型のchannelを作成
	c:=make(chan int)
	//前半と後半に分けて計算する
	go sum(a[:len(a)/2],c)
	go sum(a[len(a)/2:],c)

	x,y:=<-c,<-c	//receive from c

	fmt.Println(x,y,x+y)

}