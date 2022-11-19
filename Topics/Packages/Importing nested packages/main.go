package main

import (
	"fmt"
	"math/rand"
)

// DO NOT change the code within the main function!
func main() {
	var num int64
	fmt.Scanf("%d", &num)

	fmt.Println(rand.Int63n(num))
}
