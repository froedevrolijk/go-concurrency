package main

import (
	"fmt"
	"time"
)

func main() {
	// create a channel
	// ch := make(chan string) // unbuffured channel;
	ch := make(chan string, 1) // buffured channel;

	// start the greeter to provide a greeting
	go greet(ch)
	// sleep for a long time
	time.Sleep(5 * time.Second)
	fmt.Println("Main ready!")
	// receive greeting
	greeting := <-ch
	// sleep and print
	time.Sleep(2 * time.Second)
	fmt.Println("Greeting received!")
	fmt.Println(greeting)

}

// greet writes a greet to the given channel and then says goodbye
func greet(ch chan string) {
	fmt.Printf("Greeter ready!\n" +
		"Greeter waiting to send greeting...\n")
	// greet
	ch <- "Hello, world!"
	fmt.Println("Greeter completed!")
}

// Unbuffered (sync):

// Greeter ready!
// Greeter waiting to send greeting...

// Main ready!
// Greeter completed!

// Greeting received!
// Hello, world!

//  Buffered (async):

// Greeter ready!
// Greeter waiting to send greeting...
// Greeter completed!

// Main ready!

// Greeting received!
// Hello, world!
