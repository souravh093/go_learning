package main

import "fmt"

func add(a int, b int) { //parameter => a, b
	c := a + b
	fmt.Println(c)
}

func main() {
	add(3, 7) // argument => 3, 7
}

/*
	1. parameter vs argument
	2. First oder function
		i. standard function and named function
		ii. anonymous function
		iii. IIFE
		iv. function expression
	3. Higher order function 

	functional paradigm -> haskel, racket

	math -> logic (discrete mathematics)
	1. first order
	2. higher order

	* Logic
		1. Object (people, animal, car, etc)
		2. Property (color, student)
		3. Relation 

	# Example Logic (first order logic)
	Rule: All customer must pay their pizza bills
		1. Customer (Object)
		2. Pizza (Property)

	# Higher order logic
		Example: Any rule that applies to all customer must also apply to sudipto

	Rule example: All customer must pay tips to the waiter.const.
	
*/