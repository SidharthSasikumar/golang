package main

import (
	"fmt"
	"sync"
)

var ch = make(chan int, 5)
var ch1 = make(chan int, 5)
var wg sync.WaitGroup

func Produce1(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func Produce2(ch1 chan int) {
	for {
		val, ok := <-ch
		ch1 <- val * 2
		if ok == false {
			select {
			case msg1 := <-ch:
				ch1 <- msg1 * 2
			}
		}
	}

}

func Consume(ch1 chan int) {
	for i := 0; i < 8; i++ {
		select {
		case <-ch:
		case msg2 := <-ch1:
			fmt.Println(msg2)
		default:
			fmt.Println("")

		}
	}
	wg.Done()
}

func main() {

	go Produce1(ch)
	go Produce2(ch1)
	go Consume(ch1)
	wg.Add(1)
	wg.Wait()

}
