package controller

import "fmt"

//Cursos es un slice inicialmente vacio de cursos
var Cursos []string

//AddCurso agrega un curso
func AddCurso(nombreCurso string) {
	if ExistCurso(Cursos, nombreCurso) {
		fmt.Printf("%s ya existe\n", nombreCurso)
	} else {
		Cursos = append(Cursos, nombreCurso)
	}
}

//AgregarAlumnoACurso agrega alumno al curso
func AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {

}

//ExistCurso evalua si el curso existe
func ExistCurso(Cursos []string, nombreCurso string) bool {
	for _, curso := range Cursos {
		if curso == nombreCurso {
			return true
		}
	}
	return false
}
