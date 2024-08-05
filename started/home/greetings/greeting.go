package grettings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomMessageFormat() string {
	formats := []string {
		"Hi, %v. Welcome",
		"Great to see you, %v !",
		"Hail, %v !!! Well met!",
	}
	randomIndex := rand.Intn(len(formats))
	randomFormat := formats[randomIndex]
	return randomFormat
}
func randomErrMessageFormat() string {
	formats := []string {
		"🔥 🔥 🔥 🔥 🔥 Oops !!! :  empty name",
		"🔥 🔥 🔥 🔥 🔥 Oops !!! :  dont specific name for this...",
		"🔥 🔥 🔥 🔥 🔥 Oops !!! :  Please, input your name sir...",
	}
	randomIndex := rand.Intn(len(formats))
	randFormat := formats[randomIndex]
	return randFormat
}

// Hello returns a gretting for the name person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(randomErrMessageFormat())
	}
	// Return a gretting that embeds the name in a message.
	message := fmt.Sprintf(randomMessageFormat(), name)
	return message, nil
}