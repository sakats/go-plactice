package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // point to v
	p.X = 1e9
	(*p).Y = 1000 // 関節参照を省略しない場合
	fmt.Println(v)
}
