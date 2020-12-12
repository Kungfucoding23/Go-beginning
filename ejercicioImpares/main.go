package main

import (
	"fmt"
)

var (
	n1, n2, cont int
)

func main() {

	// Inicializo el contador
	cont = 0

	// Contador de numeros impares

	encabezado :=
		`
	****************************
	Contador de numeros impares
	****************************
	`

	//Imprimir encabezado
	fmt.Println(encabezado)
	// Solicitar el primer numero
	fmt.Println("Ingrese el primer numero")
	// Leemos el numero ingresado y lo guardamos en n1
	fmt.Scanln(&n1)
	fmt.Println("Ingrese el ultimo numero")
	fmt.Scanln(&n2)
	// Bucle tomando como inicio y fin los numeros ingresados
	for i := n1; i <= n2; i++ {
		// Evaluamos si el numero es impar
		if i%2 != 0 {
			cont++
			fmt.Printf("%d es un numero impar\n", i)
		}
	}

	resultado :=
		`
	*********
	RESULTADO
	*********
	`
	fmt.Println(resultado)
	fmt.Printf("Entre %d y %d hay %d numeros impares\n", n1, n2, cont)
}
