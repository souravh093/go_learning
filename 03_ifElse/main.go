package main

import "fmt"

func main() {
	age := 20

	if age < 18 {
		fmt.Println("You are not able to married")
	} else if age >= 18 && age < 21 {
		fmt.Println("You are able to married but not able to drink")
	} else {
		fmt.Println("You are able to married and drink")
	}
}