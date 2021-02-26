package main


import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()			//オプションを解析する
	fmt.Println((r.Form))	//サーバのプリント情報に出力される
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k,v:=range r.Form{
		fmt.Println("key: ",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello astaxie!")//ここでwに入るものがクライアントに出力される
}

func main(){
	http.HandleFunc("/",sayhelloName)		//アクセスのルーティングを設定
	err:=http.ListenAndServe(":9090",nil)	//監視するポートを設定する

	if err !=nil{
		log.Fatal("ListenAndServe: ",err)
	}

}
/*

ブラウザで下記のURLを入力
http://localhost:9090/?url_ong=111&url_long=222

ターミナルでの出力

map[url_long:[222] url_ong:[111]]
path /
scheme 
[222]
key:  url_ong
val: 111
key:  url_long
val: 222
map[]
path /favicon.ico
scheme 
[]

*/