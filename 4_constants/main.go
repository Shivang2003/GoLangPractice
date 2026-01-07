package main
import "fmt"

const b int = 10
func main() {
	const name string = "Golang"

	const (
		port = 8080
		localhost = "localhost"
	)

	fmt.Println("Constants:", b, name, port, localhost)

}