package main

import (
	"fmt"
)

func processFunction(numChan chan int) {

	for num := range numChan {
		fmt.Println("Processing this number ::", num)
	}
}

func task(done chan bool) {
	fmt.Println("processing ..........")
	defer func() { done <- true }()
}

func main() {

	done := make(chan bool)
	go task(done)

	<-done //to block

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

	// This is multiple value reading in channels

	//	}numChan := make(chan int)
	//	}go processFunction(numChan)
	//	}for {
	//	}	numChan <- rand.Intn(100)
	//	}
	// --------------------------------------------------------------------------------------------------------------------------

}
