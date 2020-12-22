package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Fib(i))
	}
}

//Fib es una funcion que devuelve la secuencia de fibonacci
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
