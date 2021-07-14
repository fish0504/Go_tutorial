package main

import (
	"fmt"
	"testing"
	//"golang.org/x/tools/go/expect"
)

//要素取得のget関数のtest
func Test01(t *testing.T) {

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
	expected1 := 1
	result1 := b.get(3, 4) //(x,y)=(3,4)

	expected2 := 2
	result2 := b.get(4, 5) //(x,y)=(4,5)

	expected3 := 3
	result3 := b.get(5, 4) //(x,y)=(5,4)
	if result1 != expected1 {
		t.Errorf("Error!\n result1 != expected1")
	} else if result2 != expected2 {
		t.Errorf("Error!\n result2 != expected2")
	} else if result3 != expected3 {
		t.Errorf("Error!\n result3 != expected3")
	}
}

func Test02(t *testing.T) {

	//get関数の動作確認ができているので
	//put関数を実行した後に
	//get関数で反映されているか確認する
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

	expected1 := 1
	b.put(3, 3, expected1)
	result1 := b.get(3, 3)

	expected2 := 7
	b.put(4, 6, expected2)
	result2 := b.get(4, 6)

	expected3 := 5
	b.put(2, 8, expected3)
	result3 := b.get(2, 8)

	//すでに数字が入っているマスに
	//数字を入れたら警告を出す
	result4:=b.put(2, 8, expected3)
	expected4:=false
	if result1 != expected1 {
		t.Errorf("Error!\n result1 != expected1")
	} else if result2 != expected2 {
		t.Errorf("Error!\n result2 != expected2")
	} else if result3 != expected3 {
		t.Errorf("Error!\n result3 != expected3")
	} else if result4 !=expected4{	
		fmt.Printf("result4=%t\n",result4)
		fmt.Printf("expected4=%t\n",expected4)

		t.Errorf("Error!\n result4 != expected4 ")
	}
}
//標準入力から座標と値を受け取って反映する
func Test03(t *testing.T){
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
	// Input (x,y):2 2
	// Input value:6
	expected:=6
	b.put(2,2,expected)
	result:=b.get(2,2)
	if expected!=result{
		t.Errorf("Error!\n result != expected")
	}
}

//
func Test04(t *testing.T){
	b := &Board{
		nums: [][]int{
			{1, 0, 3, 4, 5, 6, 7, 8, 9},
			{0, 1, 0, 0, 0, 0, 0, 0, 0},
			{3, 3, 0, 0, 0, 0, 0, 0, 0},
			{4, 4, 0, 0, 0, 0, 0, 0, 0},
			{5, 5, 0, 0, 0, 0, 0, 0, 0},
			{6, 6, 0, 0, 0, 0, 0, 0, 0},
			{7, 7, 0, 0, 0, 0, 0, 0, 0},
			{8, 8, 0, 0, 0, 0, 0, 0, 0},
			{9, 9, 0, 0, 0, 0, 0, 0, 0},
		},
		size: 9,
	}
	b2 := &Board{
		nums: [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{2, 1, 3, 4, 5, 6, 7, 8, 9},
			{3, 3, 0, 0, 0, 0, 0, 0, 0},
			{4, 4, 0, 0, 0, 0, 0, 0, 0},
			{5, 5, 0, 0, 0, 0, 0, 0, 0},
			{6, 6, 0, 0, 0, 0, 0, 0, 0},
			{7, 7, 0, 0, 0, 0, 0, 0, 0},
			{8, 8, 0, 0, 0, 0, 0, 0, 0},
			{8, 9, 0, 0, 0, 0, 0, 0, 0},
		},
		size: 9,
	}
	expected1:=true
	result1:=b.checkVer()
	
	expected2:=false
	result2:=b2.checkVer()
	if expected1!=result1{
		t.Errorf("Error!\n result(%t) != expected(%t)",result1,expected1)
	}else if expected2!=result2{
		t.Errorf("Error!\n result2(%t) != expected2(%t)",result2,expected2)
	}
}


