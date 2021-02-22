package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	//xの実際の値をvへ代入
	//v:=reflect.ValueOf(x)

	/*エラーが出る場合
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)
	*/

	//エラーが出ない場合
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)

	//vの型を出力
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

}
