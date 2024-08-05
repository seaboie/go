package grettings

import (
	"errors"
	"fmt"
)

// Hello returns a gretting for the name person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(" 🔥 🔥 🔥 🔥 🔥 Oops !!! :  empty name, try insert your name in : ")
	}
	// Return a gretting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}