package main

import "fmt"

/*
	Los MIDDLEWARE son INTERCEPTORES que permiten ejecutar instrucciones comunes a varias funciones
	que reciben y devuelven los mismos tipos de variables
	Imaginemos que tenemos un desarrollo hecho por muchas personas con muchas funciones, y ahora que queremos que a cada funcion se le adicione la grabacion en un archivo de log, seria una tarea titanica. Para eso se crearon los middlewares, que nos permiten encapsular la ejecucion de una funcion ya preexistente con otra funcion nueva, entonces lo que va a hacer es ejecutar cierto codigo previamente al llamado de la funcion que ya tenemos grabada
	Si yo voy a crear un interceptor que va a operar sobre varias funciones, esas funciones tienen que recibir y devolver el mismo tipo de variables.

	Hoy en dia los MIDDLEWARES se usan para agregarle a todas las rutinas, todas las funciones de desarrollos grandes todo lo que tiene que ver con funcionalidad, seguridad, encriptacion. Muchas operaciones como capas de un sistema con Middlewares
*/

func main() {
	fmt.Println("Inicio de programa")
	//operacionesMidd recibe dos valores, primero cual es la funcion con la que voy a trabajar y luego abriendo otros parentesis le doy los valores a la funcion sumar
	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

//El middleware recibe como parametro una funcion, donde simplemente le decimos el tipo de variables que recibe y el que devuelve esa funcion. Luego Devuelve la misma funcion ejecutando ya lo que esta en la funcion en si
func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		//Lo siguiente es lo que se aplica previo a realizar todas las funciones
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
