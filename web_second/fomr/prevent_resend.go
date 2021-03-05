package main
import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"io"
	"strconv"
	"crypto/md5"
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
//MD5(タイムスタンプ)によってクライアント側でユニークな値を割り当て
//その値をサーバに保存する
func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("method: ",r.Method)	//リクエストを受け取る方法

	if r.Method =="GET"{
		crutime :=time.Now().Unix()
		h :=md5.New()
		io.WriteString(h,strconv.FormatInt(crutime,10))
		token :=fmt.Sprintf("%x",h.Sum(nil))

		t,_:=template.ParseFiles("login.gtpl")
		t.Execute(w,token)
	}else{
		//リクエストはログインデータ
		//ログインのロジックを実行
		r.ParseForm()
		token :=r.Form.Get("token")
		if token !=""{
			//tokenの合法性を検証
		}else{
			//tokenが存在しなければエラーを出す
		}
		fmt.Println("username length:",len(r.Form["username"][0]))
		fmt.Println("username:",template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:",template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w,[]byte(r.Form.Get("username")))//クライアントに出力する


	}
	

}

func main(){

	http.HandleFunc("/",sayhelloName)
	http.HandleFunc("/login",login)

	err:=http.ListenAndServe(":9090",nil)//通信するポートの設定
	if err !=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}