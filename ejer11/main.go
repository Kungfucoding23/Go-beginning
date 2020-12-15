// Vectores y matrices
// Slices
package main

import "fmt"

var arreglo [7]int
var matriz [5][6]int
var miSlice []int

func main() {
	fmt.Println(arreglo)
	arreglo[3] = 19
	fmt.Println(arreglo)

	miVector := [10]int{23, 9, 4, 7, 5, 6, 8, 3, 2, 2}
	fmt.Println(miVector)
	fmt.Println("Contenido de miVector: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("miVector[%d] = %d\n", i, miVector[i])
	}

	matriz[2][4] = 1
	fmt.Println(matriz)
	// for i := 0; i <= 5; i++ {
	// 	for j := 0; j <= 6; j++ {
	// 		fmt.Println(matriz[i][j])
	// 	}
	// }

	miSlice = []int{3, 6, 9}
	fmt.Println(miSlice)
	miSlice = append(miSlice, 9)
	fmt.Println(miSlice)

	fmt.Println("Variante 2")
	variante2()

	fmt.Println("Variante 3")
	variante3()

	fmt.Println("Variante 4")
	variante4()

}

func variante2() {
	elementos := [5]int{2, 5, 6, 8, 9}
	porcion := elementos[2:]
	fmt.Println(elementos)
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 6, 100)
	fmt.Println(elementos)
	fmt.Printf("Largo: %d   Capacidad: %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	fmt.Println(nums)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
	fmt.Printf("Largo: %d   Capacidad: %d\n", len(nums), cap(nums))
	// go siempre toma una potencia de 2 como capacidad por eso el 16 = 2^4
	// sirve para mejorar el rendimiento, xq es facil acceso
}
