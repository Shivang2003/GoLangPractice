package main
import "fmt"

type Order struct {
	id int
	item string
	quantity int
	price float64
	status string
}

//receiver functions can be defined to add methods to structs
func ( o *Order) setStatus(newStatus string) {
	o.status = newStatus;
}

func ( o Order) getStatus() string {
	return o.status
}

func NewOrder (id int, item string, quantity int, price float64, status string) *Order{
	return &Order{
		id : id,
		item : item,
		quantity : quantity,
		price : price,
		status : status,
	}
}

func main() {
	order := Order{
		id: 1,
		item: "Laptop", //instance of an Order struct dont relate with class object
		quantity: 2,
		price: 99999.99,
	}

	order2 := NewOrder(1,"bag",1,2000.00,"On the way")

	order.status = "Shipped" //updating field
	fmt.Println("Order status: ",order.getStatus())
	order.setStatus("Confirmed")

	fmt.Println("Order:", order)
	fmt.Println("Order2:", order2)
}