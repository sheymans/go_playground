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

	fizzBuzz()

}

func double() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}

func showTen() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
