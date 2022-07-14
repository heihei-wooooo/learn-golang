package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("引数にURLを設定してください") // 例 "go run get.go jp.mercari.com"
		os.Exit(1)
	}

	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + os.Args[1]
	}

	// HTTP GETリクエストを送信する
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get()でエラー発生(%v)\n", err)
		os.Exit(1)
	}

	// ソースを取得する
	src, err := ioutil.ReadAll(resp.Body)
	// resp.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadAll()でエラー発生(%v)\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s", src)
}
