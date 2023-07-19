package main

import "fmt"

func main() {
	sum := 1
	// 保存時にフォーマッターが;を削除してしまう
	// for ; sum < 1000 ; {
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
}
