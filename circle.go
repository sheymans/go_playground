package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func area(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{
		x: 0,
		y: 0,
		r: 5,
	}
	fmt.Println(area(&c))
	fmt.Println(c.area())
}
