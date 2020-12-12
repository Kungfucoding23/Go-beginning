//Funciones
package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(2))

	num, estado := dos(6)
	fmt.Println(num, estado)
}

func uno(n int) int {
	return n * 2
}

func dos(n int) (int, bool) {
	if n == 5 {
		return n, true
	}
	return n, false
}
