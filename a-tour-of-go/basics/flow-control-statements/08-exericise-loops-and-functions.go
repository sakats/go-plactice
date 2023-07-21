package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var machine_epsilon float64 = math.Nextafter(1, 2) - 1
	z := x / 2
	for y := 0; y < 10; y++ {
		prev_z := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev_z) <= machine_epsilon {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
