package main

import "fmt"

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{2015, false, "year not divisible by 4 in common year"},
	{1970, false, "year divisible by 2, not divisible by 4 in common year"},
	{1996, true, "year divisible by 4, not divisible by 100 in leap year"},
	{2100, false, "year divisible by 100, not divisible by 400 in common year"},
	{2000, true, "year divisible by 400 in leap year"},
	{1800, false, "year divisible by 200, not divisible by 400 in common year"},
}

//EsAñoBisiesto ---
func EsAñoBisiesto(año int) bool {
	return año%4 == 0 && !(año%100 == 0 && !(año%400 == 0))
}

func main() {
	for i := range testCases {
		if EsAñoBisiesto(testCases[i].year) {
			fmt.Printf("El año %d es bisiesto\n", testCases[i].year)
		}
	}
}
