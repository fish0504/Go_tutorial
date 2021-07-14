package main

import (
	"fmt"
	
)
type Board struct{
	nums [][] int
	size int
}
func  getCkeckLen(len int)([] bool){
	alreadyExist:= make([]bool,len)
		for p:=0;p<len;p++{
		alreadyExist[p]=false
	}
	return alreadyExist
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
	
	//すでに埋まっていたらだめ
	if b.get(x,y)!=0 {
		return false;
	}else {
		x--
		y--
		b.nums[y][x]=value
		return true
	}
}
func (b* Board)input(){
	var x,y,value int
	//座標を入力
	fmt.Printf("Input (x,y):")
	fmt.Scan(&x)
	fmt.Scan(&y)

	fmt.Printf("Input value:")
	fmt.Scan(&value)
	b.put(x,y,value)
	b.show()
}
//縦方向について数字の重複がないか調べる
// 重複があればfalse,重複がなければtrue
func (b* Board)checkVer()bool{

	for k:=1;k<=b.size;k++{
		//1~9の数字が見つかっているか判定する配列
		//3が見つかっている時
		//exist[2]=true
		exist:=getCkeckLen(b.size)
		for i:=1;i<=b.size;i++{
			//まだ数字が入っていない時
			if b.get(k,i)==0{
				continue
			}else if (!exist[b.get(k,i)-1]){
				exist[b.get(k,i)-1]=true
			}else{
				return false
			}
		}
	}
	return true

}

func main(){
	b := &Board{
		nums: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 3, 0, 0, 0, 0},
			{0, 0, 0, 2, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		size: 9,
	}
	b.show()
	for ; ;{
		b.input()
		if !b.checkVer(){
			fmt.Println("failed!")
			break
		}
	}
	
}