package main

import "fmt"

func main() {
	fmt.Println("Hello world!", len("Hello"))
	fmt.Println("1 + 1 = ", 1 + 1)
	fmt.Println("1.0 + 1.0 = ", 1.0 + 1.0)
	fmt.Println("Hello"[1]) // shows 101 instead of e: it prints the byte for the character e.
	fmt.Println(true && true)

	const x string = "Hello"
	fmt.Println(x)
	y := "world"
	fmt.Println(x == y)

	double()
}

func double() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
