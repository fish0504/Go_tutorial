package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct{
	name string
	age int
}

//Stringメソッドを定義,fmt.Stringerを実装する
func (p Person)String()string{
	return "(name: "+p.name +" -age "+strconv.Itoa(p.age)+" years)"
}

func main(){
	list:=make(List,3)
	list[0]=1
	list[1]="hello"
	list[2]=Person{"Dennis",70}

	/*
	Comma-okアサーション

	Go言語の文法では、ある変数がどの型か直接判断する方法があります： value, ok = element.(T), 
	ここでvalueは変数の値を指しています。okはbool型です。
	elementはinterface変数です。Tはアサーションの型です。

	もしelementにT型の数値が存在していれば、
	okにはtrueが返されます。さもなければfalseが返ります。
	*/
	for index,element :=range list{
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
	}
		
	}
}