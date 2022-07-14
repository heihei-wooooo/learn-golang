package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("文字列を入力してください(quitで終了)")

	// 接続
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	// サーバーから受け取ったデータを表示
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	// 入力値をサーバに渡す
	var sc = bufio.NewScanner(os.Stdin)
	var txt string
	for {
		if sc.Scan() {
			txt = sc.Text()
		}
		if txt == "quit" {
			break
		}
		txt += "\n"
		_, werr := conn.Write(([]byte)(txt))
		if werr != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
	}

	conn.Close()
	<-done
}
