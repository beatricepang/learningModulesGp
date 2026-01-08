package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func M() string {
	return "Welcome to Go Module"
}

func M2() string {
	return "Welcome to Go Module 2"
}

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

// Hello returns a greeting for the named person.
func HelloWError(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// RANDOM GREETING RETURN
// random string return using math rand package
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func HelloRandom(name string) (string, error) {
	//
	if name == "" {
		return name, errors.New("empty name")
	}
	// create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil

}

// greetings for multiple people

func helloRandomMultipleSetUp(name string) (string, error) {
	// if no name is gicen return an error with a message
	if name == "" {
		return name, errors.New("empty name")
	}

	// else oit will default to:
	// message is
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func HelloRandomMultipleNames(names []string) (map[string]string, error) {
	// creates a map to associate names with messages
	messages := make(map[string]string)
	// loop through the received slice of names
	// calling the hello function to get a message for each name
	for _, name := range names {
		message, err := helloRandomMultipleSetUp(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message

	}
	return messages, nil
}
