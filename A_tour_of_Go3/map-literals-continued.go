package main

import "fmt"


type Vertex struct{
	Lat,Long float64
}


var m=map[string]Vertex{
	//型名Vertexを省略
	"bell labs":{40.5 ,-74.4},
	"Google": {34.6,-122.6},
}

func main(){
	//map[Google:{34.6 -122.6} bell labs:{40.5 -74.4}]
	fmt.Println(m)
}