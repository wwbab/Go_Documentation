package main

import (
	"fmt"

	"github.com/wwbab/Go_Documentation/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := greetings.Hello("Jisidan")
	fmt.Println(message)
}