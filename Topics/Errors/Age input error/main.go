package main

import (
	"errors"
	"fmt"
)

func checkAge(age int) error {
	switch {
	case age < 0:
		return errors.New("age can't be negative")
	case age == 0:
		return errors.New("age can't be zero")
	case age > 0 && age <= 4:
		return errors.New("babies and toddlers can't code")
	default:
		return nil
	}
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var age int
	fmt.Scanln(&age)

	err := checkAge(age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("you're old enough to learn how to code in Go")
	}
}
