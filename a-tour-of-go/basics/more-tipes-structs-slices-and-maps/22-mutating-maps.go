package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	// 2つ目はキーに値が存在するかどうかのbool
	v, ok := m["answer"]
	fmt.Println("The value:", v, " Predent?", ok)
}
