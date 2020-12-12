package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

/* Go's basic types are

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128 */

/*Zero values
Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.*/

/*Type conversions
The expression T(v) converts the value v to the type T.

Some numeric conversions:*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	tiempo()
	numAleatorio()
	raizCuadrada()
	pi()
	fmt.Println(add(2, 2))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(989))
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)

	v := 42 // change me!
	j := v
	fmt.Printf("v is of type %T\n", j)

	/*Constants
	Constants are declared like variables, but with the const keyword.

	Constants can be character, string, boolean, or numeric values.

	Constants cannot be declared using the := syntax.*/

	const World = "世界" // mundo en chino
	fmt.Println("Hello", World)
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

}

func tiempo() {
	fmt.Println("The time is ", time.Now().Unix()) // takes the current time in nanoseconds as the seed
}

func numAleatorio() {
	fmt.Println(rand.Intn(100)) // this gives you an int up to but not including 100
}

//Note: The environment in which these programs are executed is deterministic, so each time you run the example program rand.Intn will return the same number.

func raizCuadrada() {
	fmt.Println("Now you have", math.Sqrt(7), "problems.")
}

func pi() {
	fmt.Println(math.Pi)
}

func add(x, y int) int {
	return x + y
}

/*Multiple results
A function can return any number of results.

The swap function returns two strings.*/

func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

/*Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.*/

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }
