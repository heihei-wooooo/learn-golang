package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		msg := input.Text()
		go func() {
			if strings.ToLower(msg) == "hello" {
				if time.Now().Hour() < 12 {
					fmt.Fprintln(c, "[おはよう]")
				} else {
					fmt.Fprintln(c, "[おはよう]")
				}
			} else {
				fmt.Println(c, "[", strings.ToUpper(msg), "]")
			}
			time.Sleep(500 * time.Millisecond)
		}()
	}
	c.Close()
}

func main() {
	fmt.Println("サーバーstart")

	// 接続
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// 個々の接続を処理する
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
