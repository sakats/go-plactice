package main

import (
	"fmt"
	"math"
)

// xのn上をif判定の前にvとして宣言
// vがlim以下の場合にvを返却
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
