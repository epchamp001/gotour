package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	threshold := 1e-15
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("Iteration %d: z = %v\n", i, z)
		if math.Abs(z*z-x) < threshold {
			break
		}
	}
	return z
}

func main() {
	fmt.Printf("custom = %v\n", Sqrt(2))
	fmt.Printf("math   = %v\n", math.Sqrt(2))
}
