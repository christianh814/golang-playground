package main

import "fmt"

type bot interface {
	getGreeting() string
}
type enBot struct{}
type spBot struct{}


func main() {
	eb := enBot{}
	sb := spBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (enBot) getGreeting() string {
	//very custom logic
	return "Hi there!"
}

func (spBot) getGreeting() string {
	//very custom logic
	return "Hola"
}

