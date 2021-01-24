package main


import "fmt"


type Vertex struct{
	Lat,Long float64
}

var m map[string]Vertex

func main(){
	//make で　指定された型のマップを初期化する
	m=make(map[string]Vertex)
	m["Bell labs"]=Vertex{
		40.5,-2.3,
	}
	fmt.Println(m["Bell labs"])
}