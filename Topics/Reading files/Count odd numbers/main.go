package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("numbers.txt") // open 'numbers.txt' here with the os.Open function
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var oddNumberCount int // oddNumberCount will be used to count the amount of odd numbers in the file

	scanner := bufio.NewScanner(file) // create a new scanner for the file

	// Use the for scanner.Scan() loop to read the file line by line
	for scanner.Scan() {
		// DO NOT delete the code block below, it parses the read number into an integer:
		var number int
		number, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		// Write the code that checks if 'number' is odd and then increments 'oddNumberCount' below:
		if number%2 != 0 {
			oddNumberCount++
		}
	}
	fmt.Println(oddNumberCount) // print the total oddNumberCount here!
}
