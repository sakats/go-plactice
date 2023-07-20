package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x / 2
	for y := 0; y < 10; y++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
