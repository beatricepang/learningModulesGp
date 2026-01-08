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

	// var testingMessage string
	// testingMessage = greetings.Hello("Beatrice")
	// // message := greetings.Hello("Gladys")
	// fmt.Println(testingMessage)

	// // testing the error handling of the HelloWError function
	// // sets a preprended string
	// log.SetPrefix("greetings: ")
	// // when 0 is used as the variable for set flags, it indicates the date and time
	// log.SetFlags(0)

	// // requests a greeting message
	// message, err := greetings.HelloWError("")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // if no error
	// fmt.Println(message)

	// EXPLAINATION:
	// Configure the log package to print the command name ("greetings: ") at the start of its log messages, without a time stamp or source file information.
	// Assign both of the Hello return values, including the error, to variables.
	// Change the Hello argument from Gladys’s name to an empty string, so you can try out your error-handling code.
	// Look for a non-nil error value. There's no sense continuing in this case.
	// Use the functions in the standard library's log package to output error information. If you get an error, you use the log package's Fatal function to print the error and stop the program.

	// var testingMessageRandom string
	// testingMessageRandom, err := greetings.HelloRandom("gladys")
	// // random greeting testing

	// log.SetPrefix("greetings: ")
	// log.SetFlags(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(testingMessageRandom)

	// multiple random greetings testing
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.HelloRandomMultipleNames(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
