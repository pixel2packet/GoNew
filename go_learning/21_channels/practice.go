package main

import (
	"fmt"
	"time"
)

// // Sending
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

//	func sum(result chan int, num1 int, num2 int) {
//		numResult := num1 + num2
//		result <- numResult
//	}

// goroutine synchonizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("Processing...")
// 	// done <- true
// }

func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()

	// <-done // Can't send to send-only

	// emailChan <- "hello@exmaple.com" // Can't send to receive only
	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "Bonk"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}

	// emailChan := make(chan string, 5) // Queuing 100 is the limit for non blocking
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 5; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i)
	// }

	// fmt.Println("done sending")
	// close(emailChan)
	// <-done

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)

	// <-done // Blocking

	// result := make(chan int)

	// go sum(result, 4, 5)

	// res := <-result // Blocking

	// fmt.Println(res)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	// numChan <- 5

	// messageChan := make(chan string)

	// messageChan <- "ping" // Channels are blocking

	// msg := <-messageChan

	// fmt.Println(msg)
}
