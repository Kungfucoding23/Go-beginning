/*
	Quick Sort

	Sin duda, este algoritmo es uno de los más eficientes. Este método es el más rápido gracias a sus llamadas recursivas, basándose en la teoría de divide y vencerás.
	Lo que hace este algoritmo es dividir recursivamente el vector en partes iguales, indicando un elemento de inicio, fin y un pivote (o comodín) que nos permitirá segmentar nuestra lista. Una vez dividida, lo que hace, es dejar todos los mayores que el pivote a su derecha y todos los menores a su izquierda Al finalizar el algoritmo, nuestros elementos están ordenados.
	Por ejemplo, si tenemos 3 5 4 8 básicamente lo que hace el algoritmo es dividir la lista de 4 elementos en partes iguales, por un lado 3, por otro lado 4 8 y como comodín o pivote el 5. Luego pregunta, 3 es mayor o menor que el comodín? Es menor, entonces lo deja al lado izquierda Y como se acabaron los elementos de ese lado, vamos al otro lado. 4 Es mayor o menor que el pivote? Menor, entonces lo tira a su izquierda. Luego pregunta por el 8, al ser mayor lo deja donde está, quedando algo así: 3 4 5 8.
*/
package main

import "fmt"

//ArregloDesordenado es el arreglo que hay que ordenar
var ArregloDesordenado = []int{15, 3, 8, 6, 18, 1}

var n int = len(ArregloDesordenado)

//QuickSort ordena el arreglo
func QuickSort(ArregloDesordenado []int, izq, der int) []int {
	pivote := ArregloDesordenado[izq] //toma el primer elemento como pivote
	i := izq                          // i realiza la busqueda de izquierda a derecha
	j := der                          // j la realiza de derecha a izquierda
	var aux int

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
