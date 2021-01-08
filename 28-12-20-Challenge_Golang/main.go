package main

import (
	"fmt"
	"skillfactory/controller"
)

func main() {
	alumno1 := controller.NewAlumno("Alejandro", "Santin", 39850018)
	alumno2 := controller.NewAlumno("Bart", "Simpsom", 123456)
	alumno3 := controller.NewAlumno("Bob", "Esponja", 234155)
	alumno4 := controller.NewAlumno("Pepe", "Smith", 159812)
	alumno5 := controller.NewAlumno("Lisa", "Simpsom", 112342)

	//Agrego los alumnos al slice
	controller.DeleteAlumno(123456)
	controller.AddAlumno(alumno2)
	controller.AddAlumno(alumno3)
	controller.AddAlumno(alumno4)
	controller.AddAlumno(alumno5)
	controller.AddAlumno(alumno1)
	//Muestro el slice
	fmt.Println(controller.Alumnos)

	//Borro el alumno2
	fmt.Println(controller.Alumnos)
	//Lo vuelvo a borrar para verificar la validación
	controller.DeleteAlumno(123456)
	controller.AddAlumno(alumno2)

	//Agrego los cursos
	controller.AgregarCurso("React")
	controller.AgregarCurso("Node.js")
	controller.AgregarCurso("Golang")

	//Los muestro iterando el slice para que salgan uno por linea en la consola
	for i := 0; i < len(controller.Cursos); i++ {
		fmt.Println(controller.Cursos[i])
	}
	//Agrego 2 cursos que ya existen para comprobar la validación
	controller.AgregarCurso("Golang")
	controller.AgregarCurso("Node.js")
	controller.DeleteAlumno(123456)
	//Agrego alumnos a los cursos
	controller.AgregarAlumnoACurso(alumno1, "Golang")
	controller.AgregarAlumnoACurso(alumno1, "Golang")
	controller.AgregarAlumnoACurso(alumno4, "Golang")
	controller.AgregarAlumnoACurso(alumno5, "Golang")
	controller.AgregarAlumnoACurso(alumno2, "Golang")
	controller.AgregarAlumnoACurso(alumno3, "Golang")
	controller.AgregarAlumnoACurso(alumno3, "Golang")
	controller.AgregarAlumnoACurso(alumno1, "Node.js")
	controller.AgregarAlumnoACurso(alumno4, "Node.js")
	controller.AgregarAlumnoACurso(alumno3, "Node.js")
	controller.AgregarAlumnoACurso(alumno1, "React")
	controller.AgregarAlumnoACurso(alumno5, "React")
	controller.AgregarAlumnoACurso(alumno4, "React")
	controller.AgregarAlumnoACurso(alumno4, "React2")
	controller.AgregarAlumnoACurso(alumno2, "React")

	for i := 0; i < len(controller.Cursos); i++ {
		fmt.Println(controller.Cursos[i])
	}

	controller.MasAlumnos()

	controller.EnQueCursoEsta(39850018)
	controller.EnQueCursoEsta(123456)
	controller.EnQueCursoEsta(234155)

	controller.AlumnosToString(controller.Alumnos)

	fmt.Println("Arreglo desordenado: ", controller.NombreAlumnos)
	NombresOrdenados := controller.QuickSort(controller.NombreAlumnos, 0, len(controller.NombreAlumnos)-1)

	fmt.Println("Arreglo ordenado: ", NombresOrdenados)

	controller.CursosToString(controller.Cursos)
	fmt.Println("Arreglo desordenado: ", controller.NombreCursos)
	NombreCursosOrdenado := controller.QuickSort(controller.NombreCursos, 0, len(controller.NombreCursos)-1)

	fmt.Println("Arreglo ordenado: ", NombreCursosOrdenado)

	controller.CreoArchivo()

}
