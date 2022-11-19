package main

import "fmt"

func panicHandler() {
	signal := recover()

	switch signal {
	case "first":
		fmt.Println("First signal catched!")
	case "second":
		fmt.Println("Second signal handled!")
	default:
		fmt.Println("Undefined signal")
	}
}

func main() {
	defer panicHandler()

	panic("second")
}
