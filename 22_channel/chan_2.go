package main

import "fmt"

func main2() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func(chan1 chan  int, chan2 <- chan string) {
		//only send in chan (chan1 chan <-  int)
		//only receive in chan (chan2 <- chan string)
		chan1 <- 3
		fmt.Println("chan1 ", <-chan1)
		if <-chan2 == "hello" {
			fmt.Println("chan2 ", <-chan2)
		}
	}(chan1,chan2)

	go func() {
		chan2 <- "hello"
		fmt.Println("chan2 ", <-chan2)
		if <-chan1 == 3 {
			fmt.Println("chan1 ", <-chan1)
		}
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println("Received from chan1: ", msg1)
		case msg2 := <-chan2:
			fmt.Println("Received from chan2: ", msg2)
		}
	}

}
