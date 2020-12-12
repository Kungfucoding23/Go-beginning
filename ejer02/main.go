package main

import (
	"fmt"
	"strconv"
)

var numero, numAuxiliar, numTotal int
var texto string
var Estado bool = false

func main() {

	var numDelMain int
	var estado bool
	numero1, numero2, texto1 := 8, 63, "como estas"
	numero = 4
	numAuxiliar = 3
	fmt.Println(numero)
	fmt.Println(numAuxiliar)
	numTotal = numero + numAuxiliar
	fmt.Println(numTotal)
	fmt.Println(numDelMain)
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(texto1)
	texto1 = strconv.Itoa(numero)
	fmt.Println(texto1)

	estado = true
	fmt.Println(estado)
	mostrarEstado()
}

func mostrarEstado() {
	fmt.Println(Estado)
}

// hoy vimos todo lo que es instalacion (go, vsc, extensiones (go, outliner, autotest))
// se crea entorno de trabajo, una carpeta principal ...go y luego una carpeta por cada ejercicio
// definimos variables, tipos no es necesario declarar el tipo, con := lo reconoce

// la semana que viene vamos a ver ciclos, condicionales etc...
