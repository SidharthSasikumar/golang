package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
}

func main() {
	go sayHello()           // create a new goroutine
	time.Sleep(time.Second) // wait for a second
}
