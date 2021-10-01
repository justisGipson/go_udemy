// interfaces
// gonna make some chat bots

package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// ok to omit type value, and leave type
// if value not being used
func (eb englishBot) getGreeting() string {
	return "hi there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
