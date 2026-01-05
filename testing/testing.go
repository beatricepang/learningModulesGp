package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	var message string
	message = greetings.M()
	fmt.Println((message))
	message = greetings.M2()
	fmt.Println((message))

	var testingMessage string
	testingMessage = greetings.Hello("Beatrice")
	// message := greetings.Hello("Gladys")
	fmt.Println(testingMessage)

	// testing the error handling of the HelloWError function
	// sets a preprended string
	log.SetPrefix("greetings: ")
	// when 0 is used as the variable for set flags, it indicates the date and time
	log.SetFlags(0)

	// requests a greeting message
	message, err := greetings.HelloWError("")

	if err != nil {
		log.Fatal(err)
	}

	// if no error
	fmt.Println(message)
}
