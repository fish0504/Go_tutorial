package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var ret [][]uint8
	for i:=0;i<dy;i++{
		var tmp []uint8
		for k:=0;k<dx;k++{
			tmp=append(tmp,uint8(i*k))
		}
		ret=append(ret,tmp)
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
