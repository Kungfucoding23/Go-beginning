package main

import (
	"fmt"
)

//Calculo suma
var Calculo func(int, int) int = func(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Printf("Suma = %d\n", Calculo(5, 8))

	Calculo = func(n1, n2 int) int {
		return n1 - n2
	}
	fmt.Printf("Resta = %d\n", Calculo(17, 6))
}
