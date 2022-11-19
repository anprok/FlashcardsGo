package main

import "fmt"

// calculateHalf prints the result divided by 2
func calculateHalf(a int) {
	fmt.Println("The result is", a/2)
}

// halfSum sums a and b and then calls calculateHalf to print the result divided by 2
func halfSum(a int, b int) {
	calculateHalf(a + b)
}

// halfSubtract substracts b from a then calls calculateHalf to print the result divided by 2
func halfSubstract(a int, b int) {
	calculateHalf(a - b)
}

// DO NOT delete or modify the code within the main function!
func main() {
	var a, b int
	fmt.Scanln(&a, &b)
	halfSum(a, b)
	halfSubstract(a, b)
}
