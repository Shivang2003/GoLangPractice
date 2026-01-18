package main


import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count ++
		return count
	}
	
}

func main() {
	fn := counter()
	fmt.Println(fn()) // 1
	fmt.Println(fn())
}