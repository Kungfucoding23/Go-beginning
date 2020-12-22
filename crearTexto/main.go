package main

import (
	"fmt"
	"os"
)

func main() {
	createText()
	addLine()
}

func createText() {
	data, err := os.Create("./test")
	if err != nil {
		panic(err)
		return
	}
	data.Close()
}

func addLine() {
	fileName := "./test"
	if Append(fileName, "Testing...") == false {
		fmt.Println("Error al cargar la linea")
	}
}

//Append es una funcion que agrega texto al archivo
func Append(arch string, text string) bool {
	data, err := os.OpenFile(arch, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(err)
		return false
	}
	_, err1 := data.WriteString(text)
	if err1 != nil {
		fmt.Println("Hubo un error en el WriteString")
		return false
	}
	data.Close()
	return true
}
