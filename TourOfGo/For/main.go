package main

import (
	"fmt"
	"math"
)

/*For
Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// the init and post statements are optional.
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

	/*Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.*/

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i" // resultado complejo
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*If with a short statement
Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.*/

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can´t use v here, though
	return lim
}

/*If and else
Variables declared inside an if short statement are also available inside any of the else blocks.

(Both calls to pow return their results before the call to fmt.Println in main begins.)*/
