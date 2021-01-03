package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"
)

type acronymTest struct {
	input    string
	expected string
}

var stringTestCases = []acronymTest{
	{
		input:    "Portable Network 1Graphics",
		expected: "PNG",
	},
	{
		input:    "Ruby on Rails",
		expected: "ROR",
	},
	{
		input:    "First In, First Out",
		expected: "FIFO",
	},
	{
		input:    "GNU Image Manipulation Program",
		expected: "GIMP",
	},
	{
		input:    "Complementary metal-oxide semiconductor",
		expected: "CMOS",
	},
	{
		input:    "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me",
		expected: "ROTFLSHTMDCOALM",
	},
	{
		input:    "Something - I made up from thin air",
		expected: "SIMUFTA",
	},
	{
		input:    "Halley's Comet",
		expected: "HC",
	},
	{
		input:    "The Road _Not_ Taken",
		expected: "TRNT",
	},
}

// Abbreviate transform a string in its acronym.
func Abbreviate(s string) string {
	s = cleanString(s)
	parts := strings.Split(s, " ")
	sigla := ""

	for _, v := range parts {
		character := string(v[0])
		sigla += strings.ToUpper(character)
	}

	return sigla
}

func cleanString(s string) string {
	s = strings.ReplaceAll(s, " - ", " ")
	s = strings.ReplaceAll(s, "-", " ")

	reg, err := regexp.Compile("[^a-zA-Z ]+")
	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(s, "")
}

func BenchmarkAcronym(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range stringTestCases {
			Abbreviate(test.input)
		}
	}
}

func main() {
	for _, test := range stringTestCases {
		actual := Abbreviate(test.input)
		fmt.Printf("Frase original: Acronym test [%s], expected [%s], sigla [%s]\n", test.input, test.expected, actual)
	}
}
