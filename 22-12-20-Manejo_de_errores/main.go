package main

import (
	"fmt"
	// "io/ioutil"
	"log"
	"os"
)

func ejemploPanic() {
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	} else {
		fmt.Println("El programa no encontro el panic")
	}
}

func main() {
	arch := "test.txt"
	f, err := os.Open(arch)

	// defer no se ejecuta en la secuencia donde esta, se ejecuta al terminar el programa
	//defer fmt.Println("Llegamos al final del programa")
	defer func() {
		reco := recover() // si no encuentra el panic, se ejecuta el defer y el recover verifica si hubo panic, si no hubo
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	if err != nil {
		fmt.Println("Error abriendo el archivo")
		//os.Exit(1)
	}
	f.Close()

	ejemploPanic()
}

//https://steemit.com/cervantes/@orlmicron/conociendo-panic-y-recover-en-go-golang
//https://www.digitalocean.com/community/tutorials/handling-panics-in-go-es
