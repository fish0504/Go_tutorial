package main
import "fmt"

type Vertex struct{
	X,Y int
}

var (
	v1=Vertex{1,2}
	v2=Vertex{X: 3}//指定していないYは0
	v3=Vertex{}//何も初期値を指定しなければ{0,0}となる
	p=&Vertex{1,2}
)

func main(){
	fmt.Println(v1,p,v2,v3)
}