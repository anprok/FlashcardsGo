package main

import "fmt"

func makeURL() string {
	// The response parts are provided below:
	var (
		protocol   = "https://"
		domain     = "hyperskill.org/"
		path       = "golang-track?"
		parameters = "gems=1500"
	)

	// Concatenate the response parts into the 'redirectURL' variable below:
	redirectURL := protocol + domain + path + parameters

	// DO NOT remove the 'return redirectURL' statement!
	return redirectURL
}

// DO NOT MODIFY the main function below!
func main() {
	fmt.Println(makeURL())
}
