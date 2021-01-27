package main

import "fmt"

//呼び出されるたびにsumに引数xが足されていく
func adder()func(int)int{
	sum:=0
	return func(x int)int{
		sum+=x
		return sum
	}
}

/*
Go の関数はクロージャ(=それ自身の外部から変数を参照する関数値)

クロージャは参照された変数へアクセスして変えることができる
*/
func main(){

	pos,neg:=adder(),adder()
	for i:=0;i<10; i++{
		fmt.Println(
			//
			pos(i),
			neg(-2*i),
		)
	}
}