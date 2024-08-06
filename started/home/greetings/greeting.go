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
	// message := fmt.Sprintf(randomMessageFormat(), name)
	message := fmt.Sprint(randomMessageFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	if len(names) == 0 {
		return  nil, errors.New(randomErrMessageFormat())
	}
	messages := make(map[string] string)

	for _, name := range names {
		
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}