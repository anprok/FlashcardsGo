package main

import "fmt"

func main() {
	defer func() {
		if recover() != nil {
			fmt.Println("catched!")
		}
	}()
	var numbers = []int{1, 2, 3}
	var index int

	fmt.Scan(&index)

	fmt.Println(numbers[index])
}
