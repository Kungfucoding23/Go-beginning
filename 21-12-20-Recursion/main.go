//Recursion - Funcion que se llama a si misma
package main

import (
	"fmt"
)

var pot int = 3

func main() {
	potencia(1)
}

// Imprime las potencias de pot
func potencia(num int) {
	if num > 1000000 {
		return
	}
	fmt.Println(num)
	potencia(num * pot)
}
