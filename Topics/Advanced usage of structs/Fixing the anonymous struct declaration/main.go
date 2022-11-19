package main

import "fmt"

func main() {
	var grades = struct {
		Name  string
		Score int
	}{
		// initialize the fields of the anonymous 'grades' struct here
		"SpongeBob",
		0,
	}

	a := grades
	fmt.Printf("%#v", a)
}
