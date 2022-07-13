package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("数字を入力してください：")
	fmt.Scan(&n)

	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, " × ", i, " = ", n*i)
		i++
	}
}
