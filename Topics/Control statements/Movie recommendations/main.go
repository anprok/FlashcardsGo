package main

import "fmt"

func main() {
	var age int
	fmt.Scanf("%d", &age)
	switch {
	case age <= 14:
		fmt.Println("Toy Story 4")
	case age <= 18:
		fmt.Println("The Matrix")
	case age <= 25:
		fmt.Println("John Wick")
	case age <= 35:
		fmt.Println("Constantine")
	case age > 35:
		fmt.Println("Speed")
	}
}
