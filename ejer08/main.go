package main

import (
	"fmt"
)

func main() {
	fmt.Println(uno(2))

	fmt.Println("Resultado: ", Calculo(10, 10))
	fmt.Println("*********************************")
	fmt.Println("Resultado: ", Calculo(10, 10, 10, 10))
	fmt.Println("*********************************")
	fmt.Println("Resultado: ", Calculo(1, 101))
	fmt.Println("*********************************")
	fmt.Println("Resultado: ", Calculo(10, 101))
	fmt.Println("*********************************")
	fmt.Println("Resultado: ", Calculo(102, 101))
	fmt.Println("*********************************")

}

func uno(n int) (doble int) { // asignamos una variable a lo que tiene que devolver
	doble = n * 2
	return // ya sabe lo que tiene que devolver
}

// Calculo es una funcion que suma
func Calculo(n ...int) int {
	total := 0
	for _, value := range n {
		total = total + value
		//fmt.Printf("Posicion: %d  valor: %d\n", pos, value)
	}
	return total
}
