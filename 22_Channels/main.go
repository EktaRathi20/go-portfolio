package main

import (
	"fmt"
	"time"
)

/*
* If we have single goroutine then we can use channel to add waiting functionality
* If we have multiple goroutines then we can use WaitGroup to add waiting functionality
* Channels are used for communication between goroutines
 */

func printMessage(message chan string) {
	fmt.Println(<-message) // Print the message received from the channel
}

func sum(a, b int, result chan int) {
	sum := a + b
	result <- sum // Send the result to the channel
}

func task(done chan bool) {
	defer func() {
		done <- true // Signal that the task is done
	}()

	fmt.Println("Task is running")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true // Signal that the email sender is done
	}()

	for email := range emailChan {
		fmt.Println("Sending email:", email)
		time.Sleep(time.Second) // Simulate time taken to send an email
	}
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		chan1 <- "Hello from channel 1"
	}()

	go func() {
		chan2 <- "Hello from channel 2"
	}()

	// Receive messages from both channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received from chan1:", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received from chan2:", msg2)
		}
	}

	//////////////////////////////////////////

	// doneChannel := make(chan bool)
	// emailChannel := make(chan string, 100) // Buffered channel with capacity of 100

	// go emailSender(emailChannel, doneChannel)

	// for i := 1; i < 5; i++ {
	// 	emailChannel <- fmt.Sprintf("%d", i)
	// }

	// fmt.Println("done sending emails")
	// close(emailChannel) // Close the channel to signal no more emails will be sent

	// <-doneChannel

	//////////////////////////////////////////
	// messageChannel := make(chan string)
	// resultChannel := make(chan int)

	// go printMessage(messageChannel)
	// go sum(5, 10, resultChannel)
	// go task(doneChannel)

	// <-doneChannel // Wait for the task to complete

	// messageChannel <- "Hello from the main goroutine" // Send a message to the channel
	// sumResult := <-resultChannel

	// time.Sleep(1 * time.Second)

	// fmt.Println("Sum:", sumResult)

}
