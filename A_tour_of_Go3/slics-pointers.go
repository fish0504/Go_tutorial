package main
import "fmt"

//スライスはどんなデータも格納していない
//ただある配列の部分列を示している

//スライスの要素を変更すると元の配列の対応する要素が変更される

func main (){
	names:=[4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a:=names[0:2]
	b:=names[1:4]
	b[0]="lady-Gaga"//bの0番目=names[1]が変更される
	fmt.Println(a,b)
	fmt.Println(names)
}