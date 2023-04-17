package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters() {
	for i := 'a'; i <= 'j'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printNumbers()
	go printLetters()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
