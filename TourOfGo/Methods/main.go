/*Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the Abs method has a receiver of type Vertex named v.*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

// el metodo Abs() devuelve el valor absoluto de v
// v es el argumento receptor especial
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// funcion Abs() escrita como una funcion regular
func Abs1(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// otro ejemplo

type MyFloat float64

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println("Metodo:", v.Abs())
	fmt.Println("Funcion regular:", Abs1(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs2())
}
