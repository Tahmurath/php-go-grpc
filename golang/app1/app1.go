package main

import (
	"fmt"

	"log"

	//"example/app1_mod"
	"example/app2"
	//"app1_mod"
)

func main() {
	// fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("app2: ")
	log.SetFlags(0)

	// Request a greeting message.
	//message, err := app2.Hello("")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	message, err := app2.Hello("Gladys")
	message2, err2 := Hello2("dfdsd")
	// If an error was returned, print it to the console and
	// exit the program.

	if err2 != nil {
		log.Fatal(err2)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
	fmt.Println(message2)

	// Request greeting messages for the names.
	messages, err := app2.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
