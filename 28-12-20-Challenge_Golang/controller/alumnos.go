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
		fmt.Println("El alumno que desea agregar ya existe")
	} else {
		Alumnos = append(Alumnos, a)
		fmt.Println("Alumno agregado con exito")
	}
}

//DeleteAlumno borra un alumno del slice
func DeleteAlumno(dni int32) {
	existe := false
	for i := 0; i < len(Alumnos); i++ {
		if dni == Alumnos[i].Dni {
			Alumnos = append(Alumnos[:i], Alumnos[i+1:]...)
			fmt.Println("Alumno borrado")
			existe = true
		}
	}
	if !existe {
		fmt.Println("El alumno que quiere borrar no existe")
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
