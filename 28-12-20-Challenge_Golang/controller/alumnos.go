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
	if ExistAlumno(Alumnos, a) {
		fmt.Printf("El alumno %t ya existe\n", a)
	} else {
		Alumnos = append(Alumnos, a)
		fmt.Println("Alumno agregado")
	}
}

//DeleteAlumno borra un alumno del slice
func DeleteAlumno(a Alumno) {
	if ExistAlumno(Alumnos, a) {
		fmt.Println("Alumno borrado")
		// aca tengo que borrar el alumno
		Alumnos = append(Alumnos[:0], Alumnos[indexOf(a):])
	} else {
		fmt.Println("El alumno no existe")
	}
}

//ExistAlumno es una funcion que evalua si el alumno ya pertenece al arreglo
func ExistAlumno(Alumnos []Alumno, a Alumno) bool {
	for _, alumn := range Alumnos {
		if alumn == a {
			return true
		}
	}
	return false
}
