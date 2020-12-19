// quiero almacenar la asistencia de cada alumno y despues ver si un alumno estuvo presente en clase

package main

import (
	us "./alumnos"
	"fmt"
	"strings"
)

var asistencias = us.Asistencias
var todoElMes = us.TodoElMes
var asistentes = us.Asistentes
var ausentes = us.Ausentes

//Presencia es una funcion que verifica si el alumno ingresado esta presente
func Presencia() {
	var alum string
	var existe bool // Variable para evaluar si el alumno ingresado esta en el curso

	for ok := true; ok; ok = alum != "OK" { // ok para salir
		existe = false
		fmt.Println("Nombre del alumno: (Ingrese 'ok' para salir)")
		fmt.Scanln(&alum)
		alum = strings.ToUpper(alum)
		for alumno, asistencia := range asistencias {
			if alum == alumno && asistencia == 1 {
				fmt.Printf("%s esta presente\n", alum)
				existe = true

			} else if alum == alumno && asistencia == 0 {
				fmt.Printf("%s esta ausente\n", alum)
				existe = true
			}
		}
		if !existe && alum != "OK" {
			fmt.Printf("El alumno %s no pertenece al curso\n", alum)
		}
	}

}

//Lista es una funcion que muestra una lista de alumnos asistentes y otra de ausentes
func Lista() {
	// Almacena en el slice de asistentes los presentes y en el de ausentes los ausentes
	for alumno, asistencia := range asistencias {
		if asistencia == 1 {
			//fmt.Printf("%s esta presente\n", alumno)
			asistentes = append(asistentes, alumno)
		} else {
			//fmt.Printf("%s esta ausente\n", alumno)
			ausentes = append(ausentes, alumno)
		}
	}
	fmt.Println()
	fmt.Println("Alumnos presentes: ")
	for i := 0; i < len(asistentes); i++ {
		fmt.Println(asistentes[i])
	}
	fmt.Println()
	fmt.Println("Alumnos ausentes: ")
	for i := 0; i < len(ausentes); i++ {
		fmt.Println(ausentes[i])
	}
}

//Mas es una funcion que verifica que alumno tuvo mas asistencias
func Mas() {
	masAsistencias := 0
	alumnoMas := ""
	for alumno, asist := range todoElMes {
		if asist >= masAsistencias {
			masAsistencias = asist
			alumnoMas = alumno
		}
	}
	fmt.Println()
	fmt.Printf("El/la alumn@ con mas asistencias del mes es %s\n", alumnoMas)
	fmt.Println("*******************************")
}

func main() {
	Presencia()
	Lista()
	Mas()
}
