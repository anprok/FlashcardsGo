package main

import "fmt"

func main() {
	// DO NOT delete the code block below, it takes as an input 2 random strings
	var a, b string
	fmt.Scanln(&a, &b)

	word := a + b                // concatenate a and b here!
	fmt.Println(word, len(word)) // print 'word' followed by its length here!
}
