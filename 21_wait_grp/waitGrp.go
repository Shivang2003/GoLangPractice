package main

import (
	"fmt"
	"sync"
)

func routinie(id int, wg *sync.WaitGroup){
	defer wg.Done() //defer to ensure it runs after the function completes
	fmt.Println("ran id", id)
}

func main(){
	
	var wg sync.WaitGroup

	for i:=0; i<5; i++{
		go routinie(i, &wg) //pass the wait group to the routine
			wg.Add(1)
		}

		wg.Wait() //wait for all goroutines to finish
	}