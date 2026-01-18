package main

import "fmt"

func swap(a int, b int) (int, int) {
	return b, a
}

func add(a int, b int) int {
	return a + b
}

func processs() func(a, b int) int {
	return func(a, b int) int {
		return a + b
	}
}

func main() {
	x := add(2, 3)
	fmt.Println("Sum:", x)

//can return multiple values
	a, b := swap(5, 10)
	fmt.Println("After Swap:", a, b)

fn := processs()
result := fn(10, 20)
fmt.Println("Processed Result:", result)
}