package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := x
	newz := z - (z*z-x)/(2*z)

	for math.Abs(newz-z)/((newz+z)/2) > 1e-6 {
		z = newz
		newz = z - (z*z-x)/(2*z)
		fmt.Printf("guess: %v\n", z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
