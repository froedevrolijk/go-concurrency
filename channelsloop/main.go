package main

import (
	"fmt"
	"time"
)

// greetings in many languages
var greetings = []string{"Hello!", "Ciao!", "Hola!", "Hej!", "Salut!"}

func main() {
	// create a channel
	ch := make(chan string, 1)
	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")
	for greeting := range ch {
		// sleep and print
		time.Sleep(2 * time.Second)
		fmt.Println("Greeting received!", greeting)
	}

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan<- string) {
	fmt.Println("Greeter ready!")
	// greet
	for _, g := range greetings {
		ch <- g
	}
	close(ch) // close the channel to avoid a fatal deadlock. The main goroutine has no idea when the greet goroutine has finished its work and does not know when to exit from its loop
	fmt.Println("Greeter completed!")
}
