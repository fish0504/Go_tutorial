package main

import (
	"bufio"
	"net"
	"strings"

	"log"
)


func CreateListener() net.Listener {
	//Listen ポートを開く
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	return ln
}

//通信要求に応じてコネクションを開く
func ListenWorker(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	//終了後にコネクションを閉じる
	defer conn.Close()

	//受信後の処理
	for {
		//改行で終わる文字列を処理
		// will listen for message to process ending in newline (\n)
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// output message received
		log.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
func main() {

	// Create Listen port and accept to connection.
	log.Print("Launching server...")
	ln := CreateListener()

	ListenWorker(ln)

}
