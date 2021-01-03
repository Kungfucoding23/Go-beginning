package controller

import (
	"fmt"
)

//Curso ---
type Curso struct {
	Nombre  string
	Alumnos []Alumno
}

//Cursos ...
var Cursos []Curso

//AlumnosEnCurso ...
var AlumnosEnCurso []Alumno

//AgregarCurso agrega un curso
func AgregarCurso(nombreCurso string) {
	curso := Curso{Nombre: nombreCurso, Alumnos: AlumnosEnCurso}
	if ExisteCurso(Cursos, nombreCurso) {
		fmt.Println("El curso ya existe")
	} else {
		Cursos = append(Cursos, curso)
	}
}

//AgregarAlumnoACurso ---
func AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {
	for i, curso := range Cursos {
		if curso.Nombre == nombreCurso {
			if ExistAlumno(Alumnos, alumno.Dni) {
				//si el alumno no esta en el curso
				if !ExisteAlumnoEnCurso(alumno, nombreCurso) {
					Cursos[i].Alumnos = append(Cursos[i].Alumnos, alumno)
					fmt.Printf("%s agregado con exito al curso de %s\n", alumno.Nombre, nombreCurso)
				} else {
					fmt.Printf("%s ya esta en el curso %s\n", alumno.Nombre, nombreCurso)
				}
			} else {
				fmt.Printf("%s no existe\n", alumno.Nombre)
			}
		}
	}
}

//ExisteCurso evalua si existe el curso
func ExisteCurso(Cursos []Curso, nombre string) bool {
	for _, curso := range Cursos {
		if curso.Nombre == nombre {
			return true
		}
	}
	return false
}

//ExisteAlumnoEnCurso ...
func ExisteAlumnoEnCurso(alumno Alumno, nombreCurso string) bool {
	for i, curso := range Cursos {
		if curso.Nombre == nombreCurso {
			for _, nombreAlumno := range Cursos[i].Alumnos {
				if nombreAlumno == alumno {
					return true
				}
			}
		}

	}
	return false
}

//MasAlumnos recorre los cursos y devuelve el que mas alumnos tenga
func MasAlumnos() {
	aux := 0
	nombreMas := ""
	for i := range Cursos {
		if len(Cursos[i].Alumnos) > aux {
			aux = len(Cursos[i].Alumnos)
			nombreMas = Cursos[i].Nombre
		}
	}
	fmt.Printf("%s es el curso con mas alumnos, tiene %d alumnos", nombreMas, aux)
}
