package main

import "testing"



// func Test_InputNum(t *testing.T) {
// 	expected := 8
// 	//構造体
// 	b := &Board{
// 		tokens: []int{0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0,
// 			0, 0, 0, 0, 0, 0, 0, 0, 0},
// 	}
// 	//入れたいマス (2,2)
// 	//入れたい数値 8
// 	b.put(2, 2, 8)
// 	if b.get(2, 2) != expected {
// 		t.Errorf("Test_InputNum Error")
// 	}

// }

func Test_get(t *testing.T) {
	
	
	
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
	expected1:=1
	result1:=b.get(3,4)//(x,y)=(3,4)

	expected2:=2
	result2:=b.get(4,5)//(x,y)=(4,5)

	expected3:=3
	result3:=b.get(5,4)//(x,y)=(5,4)
	if (result1 != expected1){
		t.Errorf("Error!\n result1 != expected1")
	}else if (result2 !=expected2){
		t.Errorf("Error!\n result2 != expected2")
	}else if (result3 !=expected3) {
		t.Errorf("Error!\n result3 != expected3")
	}
}
