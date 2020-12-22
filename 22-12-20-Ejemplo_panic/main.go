package main

import (
	"fmt"
	"log"
	"os"
)

func ejemploPanic() {

	//abrir el archivo webserver.log para escritura
	f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// y cerrar cuando termine la funcion main
	defer f.Close()
	// asociar el manejador del archivo al log
	log.SetOutput(f)
	// agregan 10 lineas al archivo log
	for i := 0; i < 10; i++ {
		log.Printf("Error linea %v", i)
	}

	// defer func() {
	// 	reco := recover() // el tipo de dato que devuelve recover es un variant
	// 	if reco != nil {
	// 		log.Printf("Ocurrio un error que generó panic \n%v", reco) // con Fatal podemos darle un mensaje de error
	// 		//Fatalf is equivalent to Printf() followed by a call to os.Exit(1).

	// 		//Fatalf call has arguments but no formatting directivesprintf

	// 	}
	// }() //esto es la estructura de una funcion anonima que se va a ejecutar con el defer

	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}

func main() {
	ejemploPanic()
	fmt.Println("FIN")
}
