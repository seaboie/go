package main

import "fmt"
import "example.com/greetings"

func main() {
	// Get a greeting message and print it.
	message := grettings.Hello("Gladys")
	fmt.Println(message)
}