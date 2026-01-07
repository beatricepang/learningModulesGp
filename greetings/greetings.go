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

// func main() {
// 	fmt.Println(M())
// }
