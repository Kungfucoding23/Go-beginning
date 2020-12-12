/*
Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

For example, the adder function returns a closure. Each closure is bound to its own sum variable.*/

package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()  
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


/*
	pos(0) -> 0 + 0 = 0     neg(-2*0) -> 0 + 0 = 0
	pos(1) -> 0 + 1 = 1     neg(-2*1) -> 0 - 2 = -2
	pos(2) -> 1 + 2 = 3     neg(-2*2) -> -2 - 4 = -6
	...						...
	...						...
	pos(9) -> 36 + 9 = 45   neg(-2*9) -> -72 -18 = -90
*/