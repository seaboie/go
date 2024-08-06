package main

import "fmt"
import "log"
import "example.com/greetings"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin", "Kritbovorn"}
	// names := []string{}
	messages, err := grettings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(messages)
	for name, message := range messages {
		fmt.Println("-" , name, ":", message)
	}

}