package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	a := []int{5, 4, 3, 2, 1}
	a = append(a, 13)
	fmt.Println(a)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices["square"])

	delete(vertices, "square")
	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// el for tambien funciona como un while si eliminamos el i := 0 y el i++ declarando la variable afuera como un while
	fmt.Println("while:")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// podemos tambien recorrer un arreglo con un for
	fmt.Println("Recorriendo arreglo")
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index: ", index, "value: ", value)
	}

	// se puede hacer lo mismo para un mapa pero se usa "key" en lugar del index

	fmt.Println("Recorrer mapa")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key: ", key, "value: ", value)
	}

	result := sum(2, 3)
	fmt.Println(result)

	result2, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}

// creando otra funcion

func sum(x int, y int) int {
	return x + y
}

// sqrt

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}
