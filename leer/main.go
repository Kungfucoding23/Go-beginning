/*Para esto necesitamos crear un reader, que lea hasta un separador. En este caso el separador será el salto de línea; el que se escribe al presionar la tecla Enter.

Después, como la entrada de datos contiene el salto de línea que utilizamos como separador, hay que limpiarla; es en donde usamos TrimRight.*/

package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main(){
	menu :=
	`
		Bienvenido
		[1] Pizza
		[2] Tacos
		¿Que prefieres?
	`

	fmt.Print(menu)

	reader := bufio.NewReader(os.Stdin)

	entrada,_ := reader.ReadString('\n') // Leer hasta el separador de salto de linea
	eleccion := strings.TrimRight(entrada, "\r\n") // Remover el salto de linea de la entrada del usuario

	switch eleccion {
	case "1":
		fmt.Println("Prefiero Pizza")
	case "2": 
		fmt.Println("Prefiero Tacos")
	default:
		fmt.Println("Ninguno")
	}
}