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
	miCanal <- final.Sub(inicio) // devuelve la duración
}

func bucle2(miCanal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000; i++ {
	}
	final := time.Now()
	miCanal <- final.Sub(inicio) // devuelve la duración
}

// creando dos funciones bucle tardó menos tiempo
