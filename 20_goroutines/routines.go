package main

import (
	"fmt"
	"time"
)

func routinie(id int){
	fmt.Println("ran id", id)
}

func main(){
	
	for i:=0; i<5; i++{
		go routinie(i)
		go func(i int){
			fmt.Println("anonymous ran id", i)
		}(i) //anonymos function with parameter but we can also use i directly from outer scope closure
	}
time.Sleep(time.Second * 2) //to allow goroutines to complete before main exits
}