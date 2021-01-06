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
		fmt.Printf("El curso de %s ya existe\n", nombreCurso)
	} else {
		Cursos = append(Cursos, curso)
	}
}

//AgregarAlumnoACurso ---
func AgregarAlumnoACurso(alumno Alumno, nombreCurso string) {
	if ExisteCurso(Cursos, nombreCurso) {
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
	} else {
		fmt.Printf("El curso de %s no existe\n", nombreCurso)
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
	cant := 0       //Inicializo una variable cant para comparar si la cantidad de alumnos en el arreglo es mayor, si es mayor reemplaza el valor hasta terminar de recorrer todo el slice, quedando en la variable aux la cantidad mayor de alumnos
	nombreMas := "" // en nombreMas guardo el nombre del curso con mas alumnos
	for i := range Cursos {
		if len(Cursos[i].Alumnos) > cant {
			cant = len(Cursos[i].Alumnos)
			nombreMas = Cursos[i].Nombre
		}
	}
	fmt.Printf("%s es el curso con mas alumnos, tiene %d alumnos\n", nombreMas, cant)
}

//EnQueCursoEsta muestra los cursos a los que pertenece un alumno ingresado, evalua si el alumno ingresado existe, si existe recorre todos los cursos buscando el dni ingresado, al encontrarlo muestra en que curso esta. Despues tambien agregué que cuente en cuantos cursos está
func EnQueCursoEsta(dni int32) {
	cant := 0
	nombre := ""
	if ExistAlumno(Alumnos, dni) {
		for i, curso := range Cursos {
			for _, nombreAlumno := range Cursos[i].Alumnos {
				if nombreAlumno.Dni == dni {
					fmt.Printf("%s esta en el curso: %s\n", nombreAlumno.Nombre, curso.Nombre)
					cant++
					nombre = nombreAlumno.Nombre
				}
			}
		}
		fmt.Printf("%s está en %d cursos\n", nombre, cant)
	} else {
		fmt.Println("No se encontró el dni ingresado")
	}
}

//NombreCursos arreglo con los nombres de los cursos
var NombreCursos []string

//CursosToString crea un arreglo con los nombres de los cursos en string
func CursosToString(Cursos []Curso) {
	for _, curso := range Cursos {
		NombreCursos = append(NombreCursos, curso.Nombre)
	}
}
