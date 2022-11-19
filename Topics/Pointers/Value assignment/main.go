package main

import "fmt"

func main() {
	var ptr = new(int)
	var num int
	fmt.Scan(&num)

	// Write your code below:
	*ptr = num

	fmt.Println(*ptr)
}
