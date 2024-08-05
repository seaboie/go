package grettings

import "fmt"

// Hello returns a gretting for the name person
func Hello(name string) string {
	// Return a gretting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}