/*
Function values
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.*/

package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12)) // raiz cuadrada de 5^2 + 12^2 = 13

	fmt.Println(compute(hypot))    // raiz cuadrada de 3^2 + 4^2 = 5
	fmt.Println(compute(math.Pow)) // 3^4 = 81
}
