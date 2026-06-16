package main

import "fmt"

var (
	a = 50
	b = 20
)

func add(a int, b int) int {
	return a + b
}

func main(){

	z := true // This line is a Expression.

	// If Expression
	if z {
		// If Block
		fmt.Println("Hello true!")
	}

	// Invoke function
	sum := add(a, b)
	// anonymous function
	// Immediately Invoked Function Expression : IIFE
	func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}(5, 7)

	fmt.Println("add function", sum, z)
}