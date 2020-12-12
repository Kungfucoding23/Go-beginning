package main

import (
	"fmt"
)

var estado bool

func main() {
	estado = true
	numero := 10

	if estado {
		fmt.Println("estado es true")
	}
	for i := 0; i < numero; i++ {
		fmt.Println(i)
	}

	numSecreto := 9
	var num int
	var primeraVez bool = true

	for { // se repite hasta el break
		fmt.Println("Ingrese un numero: ")
		fmt.Scanln(&num)
		if num == numSecreto {
			fmt.Println("Adivinaste!")
			if primeraVez { // si es la primera vez que adivina
				primeraVez = false // cambia el valor de la variable asi la segunda vez que adivina termina el programa
				continue // continua el programa
			}else{
				break // finaliza la segunda vez que adivina
			}
		}			
	}
}
