package main

import "testing"


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
		t.Errorf("Error!\n result3 != expected3")
	}

}
