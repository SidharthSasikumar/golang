package main

import "fmt"

func main() {
	// Create a channel of type string
	messageChannel := make(chan string)

	// Start a goroutine to send a message through the channel
	go func() {
		messageChannel <- "Hello, World!"
	}()

	// Wait for a message to be sent through the channel and print it
	message := <-messageChannel
	fmt.Println(message)
}
