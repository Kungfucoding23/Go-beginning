package main

import (
	"fmt"
	"strings"
)

func main() {
	var resp string
	var asistentes []string
	var ausentes []string
	alumnos := []string{"a", "b", "c", "d"}

	for i := 0; i < len(alumnos); i++ {
		fmt.Printf("El alumno %s está presente?\n", alumnos[i])
		fmt.Scanln(&resp)
		if strings.ToUpper(resp) == "S" {
			asistentes = append(asistentes, alumnos[i])
		} else if strings.ToUpper(resp) == "N" {
			ausentes = append(ausentes, alumnos[i])
		}
	}

	fmt.Println("Lista de alumnos presentes: ")
	for i := 0; i < len(asistentes); i++ {
		fmt.Println(asistentes[i])
	}
	fmt.Println()
	fmt.Println("Lista de alumnos ausentes: ")
	for i := 0; i < len(ausentes); i++ {
		fmt.Println(ausentes[i])
	}
}
