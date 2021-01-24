package main


import "fmt"

type Vertex struct{
	Lat ,Long float64
}

var m=map[string]Vertex{
	"Bell labs":Vertex{
		40.5,-3.8,
	},
	"Google":Vertex{
		3.5,-9.8,
	},
}

func main(){
	fmt.Println(m)//		map[Bell labs:{40.5 -3.8} Google:{3.5 -9.8}]

	fmt.Println(m["Google"])//	{3.5 -9.8}
}