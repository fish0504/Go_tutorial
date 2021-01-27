package main

import "fmt"

import (
	"bufio"
	"os"
	"strings"
)
var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)
func getS() string {
	sc.Scan()
	return sc.Text()
}
func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}
func main(){
	
	s:=getS()

	if strings.Count(s,"o")+(15-len(s))>=8{
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
