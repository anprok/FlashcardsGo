package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	b = 1
	for i := 2; i <= a; i++ {
		b = b * i
	}
	fmt.Println(b)
}
