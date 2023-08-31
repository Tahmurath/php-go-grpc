package app2

import "fmt"

// func main() {
// 	fmt.Println(Hello("Hooman"))
// }

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
