package main

import "fmt"

//定数宣言は文字、文字列、boolean、数値のみ
const Pi="birth"

func main(){
	const world="世界"
	//一文字空白が開く
	fmt.Println("Hello",world,"today")
	fmt.Println("Happy",Pi,"Day")

	const Truth=true
	fmt.Println("Go rules?",Truth)
}