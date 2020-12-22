package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go slow("Alejandro")
	go slow("Lucas")
	fmt.Println("Estoy aca")
	var x string
	fmt.Scanln(&x)
}

func slow(name string) {
	letras := strings.Split(name, "")
	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}
}
