package main

import "fmt"

//ArregloDesordenado es el arreglo que hay que ordenar
var ArregloDesordenado = []string{"Santin", "Lucas", "Alejandro"}

var n int = len(ArregloDesordenado)

//QuickSort ordena el arreglo
func QuickSort(ArregloDesordenado []string, izq, der int) []string {
	pivote := ArregloDesordenado[izq] //toma el primer elemento como pivote
	i := izq                          // i realiza la busqueda de izquierda a derecha
	j := der                          // j la realiza de derecha a izquierda
	var aux string

	for i < j { // Mientras no se crucen las busquedas
		for ArregloDesordenado[i] <= pivote && i < j {
			i++
		} // busca elemento mayor que pivote
		for ArregloDesordenado[j] > pivote {
			j--
		} // busca elemento menor que pivote
		if i < j { // si no se cruzaron
			aux = ArregloDesordenado[i] // los intercambia
			ArregloDesordenado[i] = ArregloDesordenado[j]
			ArregloDesordenado[j] = aux
		}
	}
	ArregloDesordenado[izq] = ArregloDesordenado[j]
	ArregloDesordenado[j] = pivote // se coloca el pivote en su lugar de forma que tendremos los menores a su izquierda y los mayores a su derecha
	if izq < j-1 {
		QuickSort(ArregloDesordenado, izq, j-1) // se ordena el subarray izquierdo
	}
	if j+1 < der {
		QuickSort(ArregloDesordenado, j+1, der) // se ordena el subarray derecho
	}
	return ArregloDesordenado
}

func main() {
	fmt.Println(ArregloDesordenado)
	ArregloOrdenado := QuickSort(ArregloDesordenado, 0, n-1)
	fmt.Println(ArregloOrdenado)
}
