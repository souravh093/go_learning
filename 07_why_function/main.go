package main

import "fmt"

func printWelcomeMessage() {
	// Print Welcome message
	fmt.Println("Welcome to the Application")
}

func getUserName() string {
	// Get user name as input
	var name string
	fmt.Println("Enter your name -- ")
	fmt.Scanln(&name)
	return name
}

func getAgeYear() (int, int) {
	var age int
	var birthYear int
	fmt.Println("Enter your age ->")
	fmt.Scanln(&age)
	fmt.Println("Enter Date of birth year")
	fmt.Scanln(&birthYear)
	return age, birthYear
}

func resResult(name string, age int, birthYear int) {
	// Display result
	fmt.Println("Hello", name)
	fmt.Println("Your age", age, "You Date of birth year", birthYear)
	fmt.Println("Thank you for using this application")
	fmt.Println("Goodbye")
}

func main() {
	printWelcomeMessage()
	name := getUserName()
	age, birthYear := getAgeYear()
	resResult(name, age, birthYear)
}
