package main

import "fmt"

// DO NOT modify the declaration of the Employee struct!
type Employee struct {
	string
	float64
}

func main() {
	var emp Employee
	emp.string = "Donald"
	emp.float64 = 29.99

	checkAnonymousField(emp) // DO NOT delete this line!
}
