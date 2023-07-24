package main

import "fmt"

func main() {
	names := [5]string{
		"Sasuke",
		"Anse",
		"Kuro",
		"Rupinoa",
		"Sakura",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:5]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)  // deep copy のような動作をする
	fmt.Println(names) // 元となる配列の要素も変更される
}
