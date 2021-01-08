/*
	Un canal es un espacio de memoria de dialogo en el que van a dialogar distintas rutinas y cuando se aloje un valor en ese espacio de memoria, la rutina que esta pidiendo un valor a cambio va a actuar en consecuencia. Ese espacio de memoria tiene que ser de un tipo de dato, en definitica lo que va a transportar es un tipo de dato.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	canal2 := make(chan time.Duration)
	go bucle(canal1)
	go bucle2(canal2)
	// go bucle(canal1)
	// go bucle(canal1)
	fmt.Println("Llegué hasta aca")
	msg := <-canal1
	msg2 := <-canal2
	fmt.Println(msg)
	fmt.Println(msg2)
}

func bucle(miCanal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {
	}
	final := time.Now()
	//.Sub retorna la duracion
	miCanal <- final.Sub(inicio)
}

func bucle2(miCanal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {
	}
	final := time.Now()
	miCanal <- final.Sub(inicio) // devuelve la duración
}

// creando dos funciones bucle tardó menos tiempo
