package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var arch string = "./nuevoArchivo.txt"

func main() {
	//leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	//ReadFile lee el archivo, no es necesario hacer un open ni un close ni usar bufio
	datos, err := ioutil.ReadFile("./datos.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(datos))
	}
}

func leoArchivo2() {
	datos, err := os.Open("./datos.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(datos)
		for scanner.Scan() { // scaneo linea a linea
			registro := scanner.Text() // devuelve el contenido leido
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	datos.Close()
}

func graboArchivo() {
	datos, err := os.Create(arch)
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(datos, "Esta es la linea nueva")
	datos.Close()
}

func graboArchivo2() {
	fileName := arch
	if Append(fileName, "\nEsta es la segunda linea") == false {
		fmt.Println("Error al agregar la segunda linea")
	}
}

//Append es una funcion para agregar contenido al archivo
func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println("Hubo un error en el Openfile")
		return false
	}
	_, err1 := arch.WriteString(texto)
	if err1 != nil {
		fmt.Println("Hubo un error en el WriteString")
		return false
	}
	arch.Close()
	return true
}
