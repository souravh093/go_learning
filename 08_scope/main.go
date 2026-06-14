package main

import (
	"fmt"
	"example.com/mathlib"
)

// Scope block
// func main() {
// 	x := 18
// 	if x >= 18 {
// 		p := 10
// 		fmt.Println("I'm measured boy")
// 		fmt.Println("I have", p, "ten class")
// 	}

// 	add(5, 7)
// }

func main() {
	fmt.Println("Showing custom package")

	mathlib.Add(4, 7)

	mathlib.Sum(7, 9)
}
