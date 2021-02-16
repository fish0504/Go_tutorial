package main
import "fmt"

type Human struct{
	name string
	age int
	weight int
}

type Student struct{
	Human 
	speciality string
}

func main(){
	// 初期化
	mark:=Student{Human{"Mark",25,120},"Computer Science"}

	//対応するフィールドにアクセスする
	fmt.Println("His name is ", mark.name)
    fmt.Println("His age is ", mark.age)
    fmt.Println("His weight is ", mark.weight)
    fmt.Println("His speciality is ", mark.speciality)

	//フィールドの要素を修正
	mark.speciality="AI"
	fmt.Println("Mark speciality is ",mark.speciality)

	// 彼の年齢情報を修正します。
    fmt.Println("Mark become old")
    mark.age = 46
    fmt.Println("His age is", mark.age)
    // 体重情報も修正します。
    fmt.Println("Mark is not an athlet anymore")
    mark.weight += 60
    fmt.Println("His weight is", mark.weight)

	mark.Human=Human{"Marcus",55,220}
	mark.Human.age-=1

	//対応するフィールドにアクセスする
	fmt.Println("His name is ", mark.name)
    fmt.Println("His age is ", mark.age)
    fmt.Println("His weight is ", mark.weight)
    fmt.Println("His speciality is ", mark.speciality)
}