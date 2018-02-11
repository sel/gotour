package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	const eps = 0.01
	z := 1.0
	i := 0
	for d := 1.0; math.Abs(d) > eps && i < 10; i++ {
		d = (z*z - x) / (2 * z)
		z -= d
	}
	return z, i
}

func main() {
	for i := 1; i <= 1000; i++ {
		result, iterations := sqrt(float64(i))
		ae := math.Abs(result - math.Sqrt(float64(i)))
		fmt.Println(i, result, iterations, ae)
	}
}
