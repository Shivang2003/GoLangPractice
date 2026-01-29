package main

import "fmt"

type OrderStatus string

const (
	Pending   OrderStatus = "Pending"
	Shipped    = "Shipped"
	Delivered  = "Delivered"
	Cancelled  = "Cancelled"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}

func main() {
	changeStatus(Shipped)
}