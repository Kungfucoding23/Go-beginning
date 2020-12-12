package main

import (
	"fmt"
)

type Vertex struct {
	x, y int
}

// Struct Literals
/*A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.*/
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{x: 1}  // y: 0 is implicit
	v3 = Vertex{}      // x: 0 and y: 0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})
	// se puede acceder usando .
	v := Vertex{1, 2}
	v.x = 4
	fmt.Println(v.x)
	//pointers to structs
	p := &v
	p.x = 7
	fmt.Println(v)
	// Struct Literals
	fmt.Println(v1, p1, v2, v3)
}
