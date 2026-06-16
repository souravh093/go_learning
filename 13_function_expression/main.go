package main

import "fmt"

func main() {

	/* Function Expression not working
	   when you invoke in the top of
	   the function
	*/
	// add(5, 7)

	add := func(a int, b int) {
		c := a + b
		fmt.Println(c)
	}

	add(5, 7)
}
