package main

import "fmt"

func main() {
	var s = new(string) // Create a pointer variable of the 'string' type
	*s = "my string"    // Asign "my string" to the variable behind the pointer

	fmt.Println(*s) // print the variable referenced by your pointer variable here!
}
