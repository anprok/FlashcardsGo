package main

import "fmt"

func solve(ptr *int) int {
	var num = ptr // assign the value of the pointer 'ptr' to the variable 'num'

	return *num // return the value of the variable 'num'
}

// DO NOT change the contents of the main function!
func main() {
	var num = new(int)
	fmt.Scan(num)

	var test = solve(num)
	fmt.Println(test)
}
