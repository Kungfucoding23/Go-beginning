package main

import "fmt"

/*
	Los MIDDLEWARE son INTERCEPTORES que permiten ejecutar instrucciones comunes a varias funciones
	que reciben y devuelven los mismos tipos de variables
*/

func main() {
	fmt.Println("Inicio de programa")
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		//Lo siguiente es lo que se aplica previo a realizar todas las funciones
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
