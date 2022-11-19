package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	quote := readQuote() // today's quote
	file, err := os.Create("quote.txt")
	if err != nil {
		log.Fatal(err)
	}

	b, err := fmt.Fprint(file, quote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes written successfully!", b)

}

func readQuote() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	_ = scanner.Scan()
	return scanner.Text()
}
