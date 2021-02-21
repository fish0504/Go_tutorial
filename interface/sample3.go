package main

import (
	"fmt"
	"strconv"
)

type Human struct{
	name string
	age int
	phone string
}


func (h Human)Str() string{
	return "("+h.name+" - "+strconv.Itoa(h.age)+" years - "+h.phone+")"
}

func main(){
	Bob := Human{"Bob", 39, "000-7777-XXX"}

	/*
	fmtのソースファイルには
	type Stringer interface{
		String() string
	}

	と定義してあるので,Stringメソッドを持つ全ての型が
	fmt.Printlnによってコールされることができる
	*/

	//Bob の
    fmt.Println("This Human is : ", Bob)
}