package main

import (
	"fmt"
	"math/rand"
)

func processs(messageChan chan int) {
	for message := range messageChan {
		fmt.Println("Received message:", message)
	}
}

func main() {
	messageChan := make(chan int)

	go processs(messageChan) //start the processing goroutine

	for i := 0; i < 10; i++ {
		messageChan <- rand.Intn(100)
	} //send a random number to the channel

	close(messageChan)

	emailch := make(chan string, 100)

	emailch <- "Hello, World!"
	emailch <- "Welcome to Go channels!"

	fmt.Println(<-emailch)
	fmt.Println(<-emailch)
	close(emailch)
}
