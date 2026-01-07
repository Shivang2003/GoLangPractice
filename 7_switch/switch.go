package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is something else")
	}
	//no breaks needed

	//multiple conditions
	switch time.Now().Weekday() {
	case 6, 7:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//type switch

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("i is an int")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Printf("other val= %T", t) //neeed format specifier
		}
	}

	checkType(7)
	checkType("hello")
	checkType(3.14)
}
