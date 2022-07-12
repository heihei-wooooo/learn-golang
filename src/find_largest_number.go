package main

import "fmt"

func main() {
	fmt.Println("3つの数字を入力してください：")

	// 1つ目
	var first int
	fmt.Scanln(&first)
	// 2つ目
	var second int
	fmt.Scanln(&second)
	// 3つ目
	var third int
	fmt.Scanln(&third)

	if first >= second && first >= third {
		fmt.Println(first, "が最大値です")
	}
	if second >= first && second >= third {
		fmt.Println(second, "が最大値です")
	}
	if third >= first && third >= second {
		fmt.Println(third, "が最大値です")
	}
}
