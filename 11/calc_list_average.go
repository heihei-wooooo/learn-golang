package main

import "fmt"

func main() {
	var num [100]int
	var temp, sum, avg int
	fmt.Print("入力する数字の個数:")
	fmt.Scanln(&temp)
	for i := 0; i < temp; i++ {
		fmt.Print("数字を入力してください")
		fmt.Scanln(&num[i])
		sum += num[i]
	}

	avg = sum / temp
	fmt.Printf("入力された数字の個数は %d 平均は %d\n", temp, avg)
}
