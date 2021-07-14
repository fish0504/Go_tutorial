package main

import (
	"fmt"
	
)
type Board struct{
	nums [][] int
	size int
}

func outBar(len int)(){
	for i:=0;i<len+2;i++{
		fmt.Printf("--")
	}
	fmt.Println("")
	return
}
func(b* Board)show(){
	var len=b.size
	outBar(len)
	for i:=0;i<len;i++{
		
		fmt.Printf("|")
		for k:=0;k<len;k++{
			fmt.Printf("%2d",b.nums[i][k])
			if k%3==2{
				fmt.Printf("|")
			}
		}
		fmt.Println("")
		if i%3==2{
			outBar(len)
		}
	}
	return 
}

func (b *Board)	get(x,y int)int{
	x--
	y--
	return b.nums[y][x]
}
func (b *Board)	put(x,y,value int)bool{

	return true
}

func main(){
	b := &Board{
		nums: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		size: 9,
	}
	b.show()

}