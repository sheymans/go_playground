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

	averageArray()
	sliceCopy()
	mapsExcursions()
	smallest()
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

func makeArray() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func averageArray() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	var y = [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	for _, value := range y {
		total += value
	}

	fmt.Println(total / float64(len(x) + len(y)))
}

func sliceAppend() {
	slice1 := make([]int, 3, 100)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice2 := append(slice1, 4, 5)
	slice2[0] = 100 // should also change slice1
	fmt.Println(slice1, slice2)
}

func sliceCopy() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	fmt.Println(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func mapsExcursions() {
	x := make(map[int]int)
	x[1] = 10
	x[2] = 20
	value, ok := x[3]
	fmt.Println(value, ok)
}

func smallest() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
	}

	smallest := 0
	for i, value := range x {
		if i == 0 || value < smallest {
			smallest = value
		}
	}

	fmt.Println(smallest)
}