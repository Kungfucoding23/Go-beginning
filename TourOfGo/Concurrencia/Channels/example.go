package main

import (
	"fmt"
)

func main() {
	// buffered channel
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"
	c <- "three" //deadlock beacuse there are just 2 spaces
	msg := <-c
	fmt.Println(msg)
	msg = <-c
	fmt.Println(msg)
}
