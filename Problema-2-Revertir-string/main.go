package main

import (
	"fmt"
	"strings"
)

type reverseTestCase struct {
	description string
	input       string
	expected    string
}

var testCases = []reverseTestCase{
	{
		description: "an empty string",
		input:       "",
		expected:    "",
	},
	{
		description: "a word",
		input:       "robot",
		expected:    "tobor",
	},
	{
		description: "a capitalized word",
		input:       "Ramen",
		expected:    "nemaR",
	},
	{
		description: "a sentence with punctuation",
		input:       "I'm hungry!",
		expected:    "!yrgnuh m'I",
	},
	{
		description: "a palindrome",
		input:       "racecar",
		expected:    "racecar",
	},
	{
		description: "an even-sized word",
		input:       "drawer",
		expected:    "reward",
	},
}

//Reverse ...
func Reverse(s string) string {
	//Divido el string letra por letra
	var c []string = strings.Split(s, "")
	//Recorro el string invirtiendo el orden de las letras
	for i, j := 0, len(c)-1; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]
	}
	//Luego con Join uno cada letra para formar el string invertido
	return strings.Join(c, "")
}

func main() {
	for i := range testCases {
		fmt.Println(Reverse(testCases[i].input))
	}
}
