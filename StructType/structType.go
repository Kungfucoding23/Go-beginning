package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Ale", age: 24}
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println("The time is ", time.Now())
}
