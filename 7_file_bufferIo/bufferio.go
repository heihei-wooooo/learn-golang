package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fname := "./sample.txt"
	var file *os.File
	var err error

	// ファイルへの書き込み
	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開けません", fname)
	}

	var txt string
	for i := 0; i < 5; i++ {
		txt = "これはサンプルテキスト" + strconv.Itoa(i+1) + "行目\n"
		file.Write(([]byte)(txt))
	}
	// ファイルを閉じる
	file.Close()

	// ファイルからの読み込み
	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開けません", fname)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		if err := sc.Err(); err != nil {
			_ = fmt.Errorf("%sを読み込めません", fname)
			break
		}
		fmt.Println(sc.Text())
	}
}
