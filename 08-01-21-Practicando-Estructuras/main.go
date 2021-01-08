package main

import (
	"estructuras/controller"
	"fmt"
	"time"
)

func main() {
	usuario1 := controller.AltaUsuario(1, "Ing. Alejandro Santin", time.Now(), true)
	fmt.Println(usuario1)
}
