package main
import "fmt"

func name (name *string) {
	*name = "Hello, " + *name
	fmt.Println("Inside function:", *name)
}

func main() {
	nameStr := "Golang"
	fmt.Println("Before function call:", nameStr)
	name(&nameStr)
	fmt.Println("After function call:", nameStr)
}