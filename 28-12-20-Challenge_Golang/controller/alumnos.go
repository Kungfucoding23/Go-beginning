package controller

import (
	"fmt"
)

//Alumno es una estructura que contiene nombre apellido y dni
type Alumno struct {
	Nombre   string
	Apellido string
	Dni      int32
}

//Alumnos es un slice inicialmente vacio
var Alumnos []Alumno

//NewAlumno crea un nuevo alumno
func NewAlumno(nombre string, apellido string, dni int32) Alumno {
	return Alumno{
		Nombre:   nombre,
		Apellido: apellido,
		Dni:      dni,
	}
}

//AddAlumno agrega un alumno al slice
func AddAlumno(a Alumno) {
	if ExistAlumno(Alumnos, a.Dni) {
		fmt.Printf("%s ya existe\n", a.Nombre)
	} else {
		Alumnos = append(Alumnos, a)
		fmt.Printf("%s agregado con exito\n", a.Nombre)
	}
}

//DeleteAlumno borra un alumno del slice evaluando primero si existe
func DeleteAlumno(dni int32) {
	if ExistAlumno(Alumnos, dni) {
		for i := 0; i < len(Alumnos); i++ {
			if dni == Alumnos[i].Dni {
				Alumnos = append(Alumnos[:i], Alumnos[i+1:]...)
				fmt.Printf("%s fue eliminado con exito\n", Alumnos[i].Nombre)
			}
		}
	} else {
		fmt.Printf("El alumno con dni: %d no esta en el curso\n", dni)
	}
}

//ExistAlumno es una funcion que evalua si el alumno ya pertenece al arreglo
func ExistAlumno(Alumnos []Alumno, dni int32) bool {
	for _, alumn := range Alumnos {
		if alumn.Dni == dni {
			return true
		}
	}
	return false
}

//NombreAlumnos arreglo con los nombres de alumnos
var NombreAlumnos []string

//AlumnosToString crea un arreglo con los nombres de los alumnos en string
func AlumnosToString(Alumnos []Alumno) {
	for _, alumn := range Alumnos {
		NombreAlumnos = append(NombreAlumnos, alumn.Nombre)
	}
}

//QuickSort ordena alfabeticamente el arreglo de alumnos
func QuickSort(Arreglo []string, izq, der int) []string {
	pivote := Arreglo[izq] // tomo el primer alumno como pivote
	i := izq
	j := der
	var aux string

	for i < j {
		for Arreglo[i] <= pivote && i < j {
			i++
		}
		for Arreglo[j] > pivote {
			j--
		}
		if i < j {
			aux = Arreglo[i]
			Arreglo[i] = Arreglo[j]
			Arreglo[j] = aux
		}
	}
	Arreglo[izq] = Arreglo[j]
	Arreglo[j] = pivote
	if izq < j-1 {
		QuickSort(Arreglo, izq, j-1)
	}
	if j+1 < der {
		QuickSort(Arreglo, j+1, der)
	}
	return Arreglo
}
