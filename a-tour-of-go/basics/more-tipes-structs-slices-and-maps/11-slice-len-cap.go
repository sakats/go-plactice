package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4] // s[0:4]と等価. 元の配列から要素を取得.
	printSlice(s)

	// Drop its first two values.
	s = s[2:] // s[2:4]と等価.
	printSlice(s)
}

func printSlice(s []int) {
	// 容量(capacity)でスライスの最初から、元の配列の要素数をカウント
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
