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

	controller.AddAlumno(alumno1)
	controller.AddAlumno(alumno2)
	controller.AddAlumno(alumno3)
	controller.AddAlumno(alumno4)
	controller.AddAlumno(alumno5)

	fmt.Println(controller.Alumnos)

	controller.DeleteAlumno(123456)
	fmt.Println(controller.Alumnos)
	controller.DeleteAlumno(123456)

	controller.AgregarCurso("Golang")
	controller.AgregarCurso("Node.js")
	controller.AgregarCurso("React")

	for i := 0; i < len(controller.Cursos); i++ {
		fmt.Println(controller.Cursos[i])
	}
	controller.AgregarCurso("Golang")
	controller.AgregarCurso("Node.js")

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

	for i := 0; i < len(controller.Cursos); i++ {
		fmt.Println(controller.Cursos[i])
	}

	controller.MasAlumnos()

}
