package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	//"time"
)

func main(){
	db,err:=sql.Open("mysql","astaxie:@/test?charset=utf8")
	checkErr(err)

	//データの挿入
	stmt,err:=db.Prepare("INDERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res,err:=stmt.Exec("asta","research","2021-03-08")
	checkErr(err)

	id,err:=res.LastInsertId()
	checkErr(err)

	
	fmt.Println(id)
	//データの更新
	stmt,err=db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res,err=stmt.Exec("ast",id)
	checkErr(err)

	affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    //データの検索
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    //データの削除
    stmt, err = db.Prepare("delete from userinfo where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}



