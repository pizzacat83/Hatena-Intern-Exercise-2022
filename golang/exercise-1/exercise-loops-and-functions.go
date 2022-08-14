package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	newz := z - (z*z-x)/(2*z)

	for math.Abs(newz-z)/((newz+z)/2) > 1e-6 {
		z = newz
		newz = z - (z*z-x)/(2*z)
		fmt.Printf("guess: %v\n", z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
