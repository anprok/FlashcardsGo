package main

import "fmt"

func notBugButFeature() (string, string, string) {
	// Declare three variables 'protocol', 'domain', 'path'
	var (
		protocol = "https"          // put protocol here
		domain   = "hyperskill.org" // put domain here
		path     = "golang-track"   // put path to resource here
	)

	// Sending values to fixer. DO NOT EDIT the return statement!
	return protocol, domain, path
}

// DO NOT MODIFY the main function below!
func main() {
	protocol, domain, path := notBugButFeature()

	fmt.Println(protocol + "://" + domain + "/" + path)
}
