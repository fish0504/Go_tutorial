package main

import "fmt"

type Human struct{
    name string
    age int
    phone string
}

type Employee struct{
    Human //匿名フィールドHuman
    speciality string
    phone string //社員のphooneフィールド
}

func main(){
    Bob:=Employee{Human{"Bob",34,"777-444-XXXX"},"Designer","333-222"}
    fmt.Println("Bob's work phone is:",Bob.phone)

    fmt.Println("Bob's personal phone is:",Bob.Human.phone)
}