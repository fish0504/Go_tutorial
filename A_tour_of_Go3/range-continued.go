package main

import "fmt"

func main(){
	pow:=make([]int,10)

	for i:=range pow{
		pow[i]=1<<uint(i)// ==2**i
	}
	for _,value:=range pow{
		//インデックスを使わない場合
		fmt.Printf("%d\n",value)
	}
	/*
	// インデックスを使わないとエラーが出る
	for i,value:=range pow{
		//インデックスを使う場合
		fmt.Printf("%d\n",value)
	}
	*/
}

