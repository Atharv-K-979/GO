package main

import "fmt"


// type payment struct{}

// func (p payment) processPayment(amount float64) {
// 	stripePaymentGW := stripe{}
// 	stripePaymentGW.processPayment(amount)
// 	// razorpayPaymentGW := razorpay{}
// 	// razorpayPaymentGW.processPayment(amount)
// }

// type razorpay struct {}

// func (r razorpay) processPayment(amount float64) {
// 	fmt.Printf("Processing payment of amount: %.2f with Razorpay\n", amount)
// }

// now i dont want to use razorpay i want to use stripe

// type stripe struct {}

// func (s stripe) processPayment(amount float64) {
// 	fmt.Printf("Processing payment of amount: %.2f with Stripe\n", amount)
// }


// this is not good !! making new struct for every payment gateway and changing code in payment struct is not good

// here comes interface to solve this problem

type paymentGateway interface {
	processPayment(amount float64)
}

type payment struct {
	gateway paymentGateway
}

func (p payment) processPayment(amount float64) {
	p.gateway.processPayment(amount)
}

type razorpay struct {}

func (r razorpay) processPayment(amount float64) {
	fmt.Printf("Processing payment of amount: %.2f with Razorpay\n", amount)
}

type stripe struct {}

func (s stripe) processPayment(amount float64) {
	fmt.Printf("Processing payment of amount: %.2f with Stripe\n", amount)
}	


// this is dependency inversion in solid principles
func main() {
	// p := payment{}
	// p.processPayment(100.50)
	// using razorpay
	// p := payment{gateway: razorpay{}}
	// p.processPayment(100.50)

	// using stripe
	p := payment{gateway: stripe{}}
	p.processPayment(100.50)	
}