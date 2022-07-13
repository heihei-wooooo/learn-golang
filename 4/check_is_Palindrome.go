package main

import "fmt"

func main() {
	var number, remainder, temp int
	var reverse int = 0

	fmt.Print("数字を入力してください：")
	fmt.Scan(&number)

	temp = number

	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			// 入力された数字が0の場合はループ終了
			break
		}
	}

	if temp == reverse {
		fmt.Printf("%d は回分数です", temp)
	} else {
		fmt.Printf("%d は回分数ではありません", temp)
	}
}
