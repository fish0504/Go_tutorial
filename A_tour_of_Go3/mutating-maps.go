package main

import "fmt"

/*











*/

func main(){

	m:=make(map[string]int)

	//map mの操作

	//m[key]=elem
	m["Answer"]=42
	fmt.Println("The value: ",m["Anster"])

	//要素の取得
	//elem=m[key]
	m["Answer"]=48
	fmt.Println("The value:",m["Answer"])

	//要素の削除
	//delete(m,key)
	delete(m,"Answer")
	fmt.Println("The value:",m["Answer"])

	//キーに対する要素が存在するかどうかは二つ目の値で判断
	//elem,ok=m[key]
	v,ok:=m["Anster"]
	fmt.Println("The value:",v,"Present?",ok)
	
}