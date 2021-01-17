package main
import "fmt"

func main(){
	//defer:=deferへ私た関数の実行を呼び出し元の関数のおわりまで遅延させるもの
	defer fmt.Println("world")

	fmt.Println("hello")
}
//output
//hello
//world