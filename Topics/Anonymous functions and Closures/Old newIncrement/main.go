package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)

	increment := func() {
		number = number * 2 // nolint: gocritic
		number++
	}

	increment()
	fmt.Println(number)
}
