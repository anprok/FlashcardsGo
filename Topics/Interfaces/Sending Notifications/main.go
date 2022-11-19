package main

import (
	"bufio"
	"fmt"
	"os"
)

type Notification struct {
	Email   string
	Message string
}

// Finish creating the 'SendNotification()' method of the 'string' type below:
func (n Notification) SendNotification() string {
	return fmt.Sprintf("Notification sent to: %s with message: %s", n.Email, n.Message)
}

// Create the 'Notifier' interface that implements the 'SendNotification()' method below.
type Notifier interface {
	SendNotification() string
}

// Do not change the code within the main function!
// The purpose of this task is to create the 'Notifier' interface and 'SendNotification' method above!
func main() {
	var n Notifier

	var email string
	fmt.Scanln(&email)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	msg := scanner.Text()

	n = Notification{Email: email, Message: msg}
	fmt.Println(n.SendNotification())
}
