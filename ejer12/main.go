//Mapas
package main

import "fmt"

func main() {
	paises := make(map[string]string, 4)
	paises["Mexico"] = "D.F"
	paises["Brasil"] = "Brasilia"
	fmt.Println(paises)

	campeonato := map[string]int{
		"blabla": 55,
		"sdasda": 46,
		"asdasd": 3124,
		"eeasda": 23133,
	}
	for _, value := range paises {
		fmt.Println(value)
	}

	fmt.Println(campeonato)

	max := 0
	// encuentra el maximo
	for _, value := range campeonato {
		if value >= max {
			max = value
		}
	}

	fmt.Println("El puntaje mas alto es: "max)

	delete(campeonato, "sdasda")
	for equipo, puntaje := range campeonato {
		fmt.Printf("El puntaje del equipo %s es: %d\n", equipo, puntaje)
	}

	puntaje, existe := campeonato["sdasda"]

	if existe {
		fmt.Printf("El puntaje es %d y el equipo existe", puntaje)
	} else {
		fmt.Printf("El puntaje es %d y el equipo no existe", puntaje)
	}
}
