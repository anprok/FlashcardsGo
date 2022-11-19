package main

import "fmt"

func main() {
	var score int
	fmt.Scanf("%d", &score)
	if score >= 71 {
		fmt.Print("Passed!")
	} else {
		fmt.Print("Failed!")
	}
}
