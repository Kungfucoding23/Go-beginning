// Teorema: Un entero positivo n > 1 es compuesto si y solo si n tiene un divisor d que satisface que 2 <= d <= sqrt(n)
// 10000 numeros primos: http://mimosa.pntic.mec.es/jgomez53/matema/conocer/10000_primos.htm
package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
)

var (
	num, cont int
	raiz      float64
)

func main() {
	// Repite hasta que se ingrese un 0
	for ok := true; ok; ok = num != 0 {
		cont = 0
		fmt.Println("Ingrese el numero")
		fmt.Scanln(&num)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()                      //Limpia la pantalla
		raiz = math.Sqrt(float64(num)) //convierto num a float64 para poder efectuar la raiz cuadrada y guardarlo en la variable raiz float64
		//Luego convierto la variable raiz a entero dentro del for para poder compararlo con i
		for d := 2; d <= int(raiz); d++ {
			if num%d == 0 {
				fmt.Printf("%d es divisor\n", d)
				cont++
			}
		}
		if cont == 0 {
			fmt.Printf("%d es PRIMO\n", num)
		} else {
			fmt.Printf("%d es COMPUESTO\n", num)
		}
	}
}

// se me ocurrio hacer un programa que genere numeros primos
