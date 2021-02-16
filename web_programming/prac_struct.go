package main

import "fmt"


type person struct{
	name string
	age int
}
//年齢を比較して大きい方人と年齢を返す
func Older(p1,p2 person)(person,int){

	if p1.age>p2.age{
		return p1,p1.age-p2.age
	}
	
	return p2,p2.age-p1.age
	
}

func main(){
	var tom person

	//初期値を代入する
	tom.name,tom.age="Tom",18

	//bobを作成,初期化
	bob:=person{age:25,name:"Bob"}

	//structの定義の順番に従って初期化する
	//paul :=person{"Paul",43}

	tb_older,tb_diff:=Older(tom,bob)

	fmt.Printf("0f %s and %s,%s is older by %d years\n",tom.name,bob.name,tb_older.name,tb_diff)
}