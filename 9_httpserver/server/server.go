package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("サーバstart")
	// リクエスト処理
	http.HandleFunc("/", handler)
	// クライアントの受け口作成
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// リクエスト処理
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html>\n<body>\n")
	fmt.Fprintf(w, "<h1>テストテスト</h1>\n")
	fmt.Fprintf(w, "<p>サーバは：%q</p>\n", r.Host)
	fmt.Fprintf(w, "<p>アドレスは：%q</p>\n", r.RemoteAddr)
	fmt.Fprintf(w, "</body>\n</html>\n")
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
}
