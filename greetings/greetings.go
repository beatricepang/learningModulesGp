package greetings

import "fmt"

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

// func main() {
// 	fmt.Println(M())
// }
