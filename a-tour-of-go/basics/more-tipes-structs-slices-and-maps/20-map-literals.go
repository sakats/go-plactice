package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 宣言＋初期化パターン
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 宣言のみ
var x map[string]Vertex

func main() {
	fmt.Println(m)

	// makeメソッドで初期化
	x = make(map[string]Vertex)
	x["My Home"] = Vertex{
		39.01621, 139.90541,
	}
	fmt.Println(x)
}
