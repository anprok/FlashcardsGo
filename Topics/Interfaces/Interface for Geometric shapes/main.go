package main

import (
	"fmt"
	"math"
)

// Do not change the type declarations below!
type (
	Square struct {
		Side float64
	}

	Circle struct {
		Radius float64
	}

	Shape interface {
		Area() float64
	}
)

// DO NOT change these lines -- they create the proper output string:
func (s Square) String() string { return fmt.Sprintf("Square area: %.2f", s.Area()) }
func (c Circle) String() string { return fmt.Sprintf("Circle area: %.2f", c.Area()) }

// Implement the interface methods for the 'Square' and 'Circle' structs below:
func (s Square) Area() float64 { return s.Side * s.Side }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func main() {
	// Do NOT change the code below! This reads the input and outputs as required.
	var length float64
	fmt.Scanln(&length)

	for _, shape := range []Shape{Square{length}, Circle{length / 2}} {
		fmt.Println(shape)
	}
}
