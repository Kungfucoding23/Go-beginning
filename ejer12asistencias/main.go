// guardar info del curso por ejemplo con la assitencia
// quiero almacenar la asistencia de casa alumno y despues ver si un alumno estuvo presente en una clase

package main

import "fmt"

func main() {
	asistencias := map[string]int{
		"Alumno1": 1,
		"Alumno2": 0,
		"Alumno3": 1,
		"Alumno4": 1,
	}
	for alumno, asistencia := range asistencias {
		if asistencia == 1 {
			fmt.Printf("El alumno %s esta presente\n", alumno)
		} else {
			fmt.Printf("El alumno %s esta ausente\n", alumno)
		}

	}
	fmt.Println()

	var alum string
	fmt.Println("Ingrese el alumno")
	fmt.Scanln(&alum)

	for al, asis := range asistencias {
		if alum == al && asis == 1 {
			fmt.Printf("El alumno %s esta presente\n", alum)
		} else if alum == al && asis == 0 {
			fmt.Printf("El alumno %s esta ausente\n", alum)
		} else {
			fmt.Println("El alumno no existe")
			break
		}
	}
	// alumnos := []string{"uno", "dos", "tres", "cuatro", "cinco"}
	// asistencia1 := []int{1, 0, 1, 1, 0}

	// for i := 0; i < 5; i++ {
	// 	if asistencia1[i] == 1 {
	// 		fmt.Printf("El alumno %s esta presente\n", alumnos[i])
	// 	} else {
	// 		fmt.Printf("El alumno %s esta ausente\n", alumnos[i])
	// 	}
	// }

	todoElMes := map[string]int{
		"Alumno 1": 20,
		"Alumno 2": 16,
		"Alumno 3": 21,
		"Alumno 4": 19,
	}

	masAsistencias := 0
	alumnoMas := ""
	for alumno1, asist := range todoElMes {
		if asist >= masAsistencias {
			masAsistencias = asist
			alumnoMas = alumno1
		}
	}
	fmt.Printf("El alumno con mas asistencias del mes es el %s", alumnoMas)
}
