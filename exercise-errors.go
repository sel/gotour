package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot take square root of %g!", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	const eps = 0.01
	z, d := 1.0, 1.0
	for i := 0; math.Abs(d) > eps && i < 10; i++ {
		d = (z*z - x) / (2 * z)
		z -= d
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
