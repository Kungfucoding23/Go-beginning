package main

import (
	us "./user"
	"fmt"
	"time"
)

type alumno struct {
	us.Usuario
}

func main() {
	user := new(alumno)
	user2 := new(alumno)
	user.AltaUsuario(1, "Bob Esponja", time.Now(), false)
	user2.AltaUsuario(2, "Patricio", time.Now(), true)
	fmt.Println(user.Usuario.Nombre)
	if user.Usuario.Status {
		fmt.Println("Presente")
	} else {
		fmt.Println("Ausente")
	}
	fmt.Println(user2.Usuario.Nombre)
}
