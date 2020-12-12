package main

import (
	"fmt"
	"bufio" // para entrada y salida (scanear o imprimir)
	"os"
)

var n1, n2, result int
var leyenda string

func main(){
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&n1)
	fmt.Println("Ingrese un numero")
	fmt.Scanln(&n2)
	fmt.Println("Descripcion: ")

	escanerAux := bufio.NewScanner(os.Stdin) // standard input de mi sistema operativo
	if escanerAux.Scan(){
		leyenda = escanerAux.Text()
	}
	result = n1 + n2
	fmt.Println(leyenda, result)
}