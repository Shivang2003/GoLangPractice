package main

import "fmt"

func main() {

	randomNumber := 7

	if num := randomNumber; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

}