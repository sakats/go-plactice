package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)

	s = s[1:4]
	fmt.Println(s)
	fmt.Printf("スライスの長さは %d\n", len(s))

	s = s[:] // [0:3]と等価(下限が0,上限がスライス長)
	fmt.Println(s)
	fmt.Printf("スライスの長さは %d\n", len(s))

	s = s[:2] // [0:2]と等価
	fmt.Println(s)
	fmt.Printf("スライスの長さは %d\n", len(s))

	s = s[1:] // [1:2]と等価
	fmt.Println(s)
	fmt.Printf("スライスの長さは %d\n", len(s))
}
