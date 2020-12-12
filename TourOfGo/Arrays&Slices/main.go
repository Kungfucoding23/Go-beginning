package main

import (
	"fmt"
	"strings"
)

func main() {
	//The type [n]T is an array of n values of type T.
	var a [2]string
	a[0] = "Hello"
	a[1] = "world!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/*An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
	The type []T is a slice with elements of type T.*/

	var s []int = primes[1:4] //This selects a half-open range which includes the first element, but excludes the last one.
	fmt.Println(s)

	/*Slices are like references to arrays
	A slice does not store any data, it just describes a section of an underlying array.

	Changing the elements of a slice modifies the corresponding elements of its underlying array.

	Other slices that share the same underlying array will see those changes.*/

	names := [4]string{
		"Jhon",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	b[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)

	/*Slice literals
	A slice literal is like an array literal without the length.

	This is an array literal:
	[3]bool{true, true, false}

	And this creates the same array as above, then builds a slice that references it:
	[]bool{true, true, false}*/

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	k := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(k)

	/*Slice defaults
	When slicing, you may omit the high or low bounds to use their defaults instead.

	The default is zero for the low bound and the length of the slice for the high bound.

	For the array

	var a [10]int
	these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]*/

	u := []int{2, 3, 5, 7, 11, 13}

	u = u[1:4]
	fmt.Println(u)

	u = u[:2]
	fmt.Println(u)

	u = u[1:]
	fmt.Println(u)

	/*Slice length and capacity
	A slice has both a length and a capacity.

	The length of a slice is the number of elements it contains.

	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

	You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.*/

	t := []int{2, 3, 5, 7, 11, 13}
	printSlice(t)

	// Slice the slice to give it zero length
	t = t[:0]
	printSlice(t)

	// Extend its length
	t = t[:4]
	printSlice(t)

	// Drop its first two values
	t = t[2:]
	printSlice(t)

	/*
		Nil slices
		The zero value of a slice is nil.

		A nil slice has a length and capacity of 0 and has no underlying array.*/

	var o []int
	fmt.Println(o, len(o), cap(o))
	if o == nil {
		fmt.Println("nil!")
	}

	/*Creating a slice with make
	Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5)  // len(a)=5
	To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4*/

	/*
		Slices of slices
		Slices can contain any type, including other slices.*/

	// Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var g []int
	printSlice(g)

	// Append works on nil slices
	g = append(g, 0)
	printSlice(g)

	// The slice grows as needed
	g = append(g, 1)
	printSlice(g)

	// We can add more than one element at a time
	g = append(g, 2, 3, 4, 5)
	printSlice(g)

	/*Range
	The range form of the for loop iterates over a slice or map.

	When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.*/

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	/*Range continued
	You can skip the index or value by assigning to _.

	for i, _ := range pow
	for _, value := range pow
	If you only want the index, you can omit the second variable.

	for i := range pow*/

	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i for example 1 << 5 is 1 times 2, 5 times.
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(t []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(t), cap(t), t)
}
