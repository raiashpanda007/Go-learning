package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("Sending email to ...", email)
		time.Sleep(time.Second)
	}

}
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
	emailChan := make(chan string, 100)
	go emailSender(emailChan, done)
	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Completed Sending")
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
