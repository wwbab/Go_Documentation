package main

import (
	"fmt"
	"log"

	"github.com/wwbab/Go_Documentation/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Wuwenbin","Jisidan","Cat"}

	// Request greeting messages for names.
	messages, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of messages
	// to the console.
	fmt.Println(messages)

	// Request a greeting message for name.
	message, err := greetings.Hello("Jisidan")
	// If an error was returned, print it to the console and 
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	//to the console.
	fmt.Println(message)
}