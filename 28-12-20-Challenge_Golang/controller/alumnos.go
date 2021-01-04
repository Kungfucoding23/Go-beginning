package controller

import "fmt"

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
