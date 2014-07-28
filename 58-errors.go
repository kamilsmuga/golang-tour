package main

import (
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, ErrNegativeSqrt(f)
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z = z - (math.Sqrt(z)-f)/2*z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number"
}
