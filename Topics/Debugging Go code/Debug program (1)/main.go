package main

import (
	"fmt"
)

func main() {
	var prev, current = 0, 1
	var sum int
	fmt.Println(current)
	for i := 0; i < 9; i++ {
		sum = fibonacci(prev, current)
		fmt.Println(sum)
		prev = current
		current = sum
	}
}

func fibonacci(prev, current int) int {
	return prev + current
}
