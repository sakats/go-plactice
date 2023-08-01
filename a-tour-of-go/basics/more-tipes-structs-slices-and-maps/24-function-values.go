package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 関数を変数に代入(関数値)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// 関数値を引数や戻り値に利用可能
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
