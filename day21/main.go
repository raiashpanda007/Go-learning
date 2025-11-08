package main

import (
	"fmt"
	"math/rand"
)

func processFunction(numChan chan int) {

	for num := range numChan {
		fmt.Println("Processing this number ::", num)
	}
}
func main() {
	//	messageChan := make(chan string)
	// messageChan <- "ping" // This is to send data to channels.
	//	msg := <-messageChan

	//	fmt.Println(msg)

	// This creates deadlock as channels block the thread.

	// ---------------------------------------------------------------------------------------------------------
	// This one for single number.
	// numChan := make(chan int)

	// go processFunction(numChan)
	// numChan <- 5
	// -------------------------------------------------------------------------------------------------------------
	numChan := make(chan int)
	go processFunction(numChan)
	for {
		numChan <- rand.Intn(100)
	}

}
