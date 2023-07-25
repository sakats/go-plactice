package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // 長さ･容量が0
	fmt.Printf("%v\n", s)
	if s == nil {
		fmt.Println("s is nil!")
	}
}
