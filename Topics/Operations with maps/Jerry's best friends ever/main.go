package main

import (
	"fmt"
	"log"
)

func createSet() map[string]bool {
	// fix the type of the bestFriends map values.
	bestFriends := map[string]bool{
		"George Costanza": true,
		"Elaine Benes":    true,
		"Cosmo Kramer":    true,
	}

	return bestFriends // do not delete this line!
}

// DO NOT delete or modify the contents of the main function!
func main() {
	mySlice := make([]string, 2)
	fmt.Scanln(&mySlice[0], &mySlice[1])

	bfSet := createSet()

	if !bfSet[mySlice[0]+" "+mySlice[1]] {
		log.Fatal("bool type sets should have 'true' as map values!")
	}
	fmt.Println(mySlice[0], mySlice[1], "is one of Jerry's best friends")
}
