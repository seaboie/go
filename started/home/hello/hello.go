package main

import "fmt"
import "log"
import "example.com/greetings"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := grettings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}