package main

import "fmt"

var arr [5]int

func main() {
	fmt.Println(arr)

	for _, value := range arr {
		fmt.Println(value)
	}
}

// func orden(arr[5] int) int{
// 	for i := 0; i < 5; i++{

// 	}
// }
