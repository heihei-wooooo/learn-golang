package main

import (
	"fmt"
	"os"
)

func main() {
	fname := "sample.txt"
	var file *os.File
	var err error

	// ファイルへの書き込み
	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開けません", fname)
	}

	txt := "これはサンプルテキストです"
	file.Write(([]byte)(txt))
	// ファイルを閉じる
	file.Close()

	// 読み込みバッファーサイズ
	const BUFSIZE = 1024

	// ファイルからの読み込み
	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開ません", fname)
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			_ = fmt.Errorf("%sの読み込みでエラーが発生しましました", fname)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
