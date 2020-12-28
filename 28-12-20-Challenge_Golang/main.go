package main

import (
	"fmt"
	"skillfactory/controller"
)

func main() {
	alumno1 := controller.NewAlumno("Alejandro", "Santin", 39850018)
	alumno2 := controller.NewAlumno("Alejandro", "Santin", 39850018)
	alumno3 := controller.NewAlumno("Bob", "", 123456)

	fmt.Println(alumno1)

	controller.AddAlumno(alumno1)
	controller.AddAlumno(alumno2)
	controller.AddAlumno(alumno3)

	fmt.Println(controller.Alumnos)
	//controller.AddAlumno(alumno1)

	controller.AddCurso("Curso de Go")
	controller.AddCurso("Curso de Go")
	controller.AddCurso("Pasear perros")

	fmt.Println(controller.Cursos)

}
