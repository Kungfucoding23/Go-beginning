/*
Problema 4: Era espacial
https://exercism.io/tracks/go/exercises/space-age

Dada una edad en segundos, calcule la edad de alguien en términos de años solares de un planeta dado.

Mercurio: período orbital 0,2408467 años terrestres
Venus: período orbital 0,61519726 años terrestres
Tierra: período orbital 1.0 años terrestres, 365.25 días terrestres o 31557600 segundos
Marte: período orbital 1.8808158 años terrestres
Júpiter: período orbital 11,862615 años terrestres
Saturno: período orbital 29.447498 años terrestres
Urano: período orbital 84.016846 años terrestres
Neptuno: período orbital 164.79132 años terrestres

Entonces, si le dijeran que alguien tenía 1,000,000,000 de segundos de edad, debería poder decir que tiene 31.69 años terrestres.
*/
package main

import "fmt"

var testCases = []struct {
	description string
	planet      Planet
	seconds     float64
	expected    float64
}{
	{
		description: "age on Earth",
		planet:      "Earth",
		seconds:     1000000000,
		expected:    31.69,
	},
	{
		description: "age on Mercury",
		planet:      "Mercury",
		seconds:     2134835688,
		expected:    280.88,
	},
	{
		description: "age on Venus",
		planet:      "Venus",
		seconds:     189839836,
		expected:    9.78,
	},
	{
		description: "age on Mars",
		planet:      "Mars",
		seconds:     2129871239,
		expected:    35.88,
	},
	{
		description: "age on Jupiter",
		planet:      "Jupiter",
		seconds:     901876382,
		expected:    2.41,
	},
	{
		description: "age on Saturn",
		planet:      "Saturn",
		seconds:     2000000000,
		expected:    2.15,
	},
	{
		description: "age on Uranus",
		planet:      "Uranus",
		seconds:     1210123456,
		expected:    0.46,
	},
	{
		description: "age on Neptune",
		planet:      "Neptune",
		seconds:     1821023456,
		expected:    0.35,
	},
}

//Planet type.
type Planet string

// earthSec is how many seconds in a year on earth
var earthSec float64 = 31557600

// secInAYear is how many seconds in a year on a planet
var secInAYear = map[Planet]float64{
	"Mercury": 0.2408467 * earthSec,
	"Venus":   0.61519726 * earthSec,
	"Earth":   1 * earthSec,
	"Mars":    1.8808158 * earthSec,
	"Jupiter": 11.862615 * earthSec,
	"Saturn":  29.447498 * earthSec,
	"Uranus":  84.016846 * earthSec,
	"Neptune": 164.79132 * earthSec,
}

// Age returns how many years old on given planet. If planet is not defined, it returns -1.
func Age(t float64, planet Planet) float64 {
	k, ok := secInAYear[planet]
	if ok {
		return t / k
	}
	return -1
}

func main() {
	for i := range testCases {
		// fmt.Println(testCases[i].seconds, " del planeta ", testCases[i].planet, " son ", Age(testCases[i].seconds, testCases[i].planet), " años terrestres")
		fmt.Printf("Tiene %f años terrestres\n", Age(testCases[i].seconds, testCases[i].planet))
	}
}
