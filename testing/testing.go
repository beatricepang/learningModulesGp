package main

import (
	"example/greetings"
	"fmt"
)

func main() {
	var message string
	message = greetings.M()
	fmt.Println((message))
	message = greetings.M2()
	fmt.Println((message))
}
