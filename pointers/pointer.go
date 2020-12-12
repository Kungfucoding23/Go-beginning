// Pointers
//Go has pointers. A pointer holds the memory address of a value.
//The & operator generates a pointer to its operand.
//The type *T is a pointer to a T value. Its zero value is nil.
package main

import (
	"fmt"
)

func main() {
	i := 7
	inc(&i)
	fmt.Println(i)
}

func inc(x *int) {
	*x++
}
