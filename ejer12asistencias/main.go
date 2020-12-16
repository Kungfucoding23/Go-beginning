// guardar info del curso por ejemplo con la assitencia
// quiero almacenar la asistencia de casa alumno y despues ver si un alumno estuvo presente en una clase

package main

import "fmt"

func main() {
	asistencias := map[string]int{
		"Homero": 1,
		"Bart":   0,
		"Lisa":   1,
		"Marge":  1,
		"Maggie": 1,
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

	// for alumno, asistencia := range asistencias {
	// 	fmt.Println("Ingrese el alumno")
	// 	fmt.Scanln(&alum)
	// 	if alum == alumno && asistencia == 1 {
	// 		fmt.Printf("El alumno %s esta presente\n", alum)

	// 	} else if alum != alumno {
	// 		fmt.Printf("El alumno %s no pertenece al curso\n", alum)

	// 	} else if alum == alumno && asistencia == 0 {
	// 		fmt.Printf("El alumno %s esta ausente\n", alum)

	// 	}
	// }

	for ok := true; ok; ok = alum != "ok" { // ok para salir
		fmt.Println("Ingrese el alumno")
		fmt.Scanln(&alum)
		for alumno, asistencia := range asistencias {
			if alum == alumno && asistencia == 1 {
				fmt.Printf("El alumno %s esta presente\n", alum)

				// } else if alum != alumno {
				// 	fmt.Printf("El alumno %s no pertenece al curso\n", alum)

			} else if alum == alumno && asistencia == 0 {
				fmt.Printf("El alumno %s esta ausente\n", alum)

			}
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
		"Homero": 20,
		"Bart":   16,
		"Lisa":   21,
		"Marge":  19,
		"Maggie": 20,
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
