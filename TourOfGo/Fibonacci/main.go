package main

import "fmt"

var a []int

// Funcion clausura
func fibonacci() func(int) int {
	// Recibe un valor x que es el contador del for de la funcion main
	return func(x int) int {
		// si es 2 o mayor, sumará los dos valores anteriores guardados en el slice a, guardará este valor con la función append(a, xf) y finalmente lo retornará
		if x >= 2 {
			xf := a[x-1] + a[x-2]
			a = append(a, xf)
			return xf
		}
		// si x es 0 o 1, lo guardará en el slice global a y lo retornará
		a = append(a, x)
		return (a[x])
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Print(f(i), " ")
	}
}
