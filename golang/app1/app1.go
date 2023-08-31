package main

import (
	"fmt"

	"rsc.io/quote"

	"example/app2"
)

func main() {
	fmt.Println(quote.Go())

	message := app2.Hello("Gladys")
	fmt.Println(message)
}
