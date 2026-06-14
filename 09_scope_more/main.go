package main

import "fmt"

var (
	a = 20
	b = 30
)

func printNum(num int) {
	fmt.Println(num)
}

func add(num1 int, num2 int) {
	res := num1 + num2
	printNum(res)

}

func main() {
	add(a, b)
}
