package main
import "fmt"

type paymenter interface {
	makePayment(amount float64)
}

type payment struct {
	gateway paymenter
}

type paypal struct{}

func (p paypal) makePayment(amount float64){
	fmt.Println("Payment of", amount, "made using PayPal")
}

type razorPay struct{}

func (r razorPay) makePayment(amount float64){
//logic for razorpay payment
fmt.Println("Payment of", amount, "made using RazorPay")
}

func main(){
	newPayment := payment{}
	newPayment.gateway = paypal{}
	newPayment.gateway.makePayment(1000.00)
}