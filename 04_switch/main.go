package main

import "fmt"

func main() {
	a := 5

	switch a {
	case 1:
		fmt.Println("Hello", a)
	case 2:
		fmt.Println("Hello", a)
	default:
		fmt.Println("Hello Default")
	}
}
