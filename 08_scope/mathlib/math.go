package mathlib

import "fmt"

func Add(x int, y int) {
	z := Sum(x, y)
	fmt.Println(z)
}

func Sum(x int, y int) int {
	return x + y
}
