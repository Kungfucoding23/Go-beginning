package main

import (
	"fmt"
)

/*Cuando creamos una función y le pasamos una variable como argumento, lo que hace la función es hacer una copia del valor de la variable y trabajar con ese valor, por lo que la variable que pasamos como argumento no se modifica.*/
func increase(v int) int {
	v++
	return v
}

/*Increase recibe un puntero de tipo entero*/
func Increase(v *int) int {
	/*Desreferenciamos la variable v para obtener su valor e incrementarlo en 1 */
	*v++
	return *v
}

func main() {
	v := 19
	increase(v)
	fmt.Println("El valor de v es: ", v)
	/*Podemos ver entonces que el valor de la variable v no se incrementa y sigue siendo 19, ya que la función Increase hace una copia de la variable v y realiza el incremento sobre la copia y no sobre la variable.*/
	fmt.Printf("La direccion de memoria de v es: %d\n", &v)
	/*La funcion Increase recibe un puntero, utilizamos el operador de direccion & para pasar la direccion de memoria de v */
	Increase(&v)
	fmt.Println("El valor de v ahora es: ", v)

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
