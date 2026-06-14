package main

import "fmt"

func add(numb1 int, numb2 int) int {
	sum := numb1 + numb2

	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2

	mul := num1 * num2

	return sum, mul
}

func sayHello (name string) {
	fmt.Println("Welcome to the go course,",name)
}

func main() {
	// a := 20
	// b := 50

	// total := add(a, b)

	// fmt.Println(total)

	// sum, mul := getNumbers(a, b)

	// fmt.Println(sum, mul)	

	sayHello("Sourav Halder")
}