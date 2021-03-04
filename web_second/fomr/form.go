package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm() //urlが渡すオプションを解析


	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)

	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v:=range r.Form{
		fmt.Println("Key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie!")		//ここでwに書き込まれたものがクライアントに出力される
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Println("method:",r.Method)		//リクエストを取得するメソッド

	if r.Method =="GET"{
		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,nil)
	}else{
		//formの内容を解析
		//r.Parseformがなければフォームのデータに対して操作を行うことはできない
		r.ParseForm()
		//ログインデータがリクエストされ,ログインのロジック判断が実行される
		fmt.Println("username:",r.Form["username"])
		fmt.Println("password:",r.Form["password"])
		//r.Formには全てのデータが含まれる
		fmt.Println("everything",r.Form)
	}
}

func main(){
	http.HandleFunc("/",sayhelloName)		//アクセスのルーティングを設定
	http.HandleFunc("/login", login)		//アクセスのルーティングを設定

	err:=http.ListenAndServe(":9090",nil)	//監視するポートを設定
	if err !=nil{
		log.Fatal("ListenAndServe: ",err)
	}

}