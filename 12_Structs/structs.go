package main

import (
	"fmt"
	"time"
)

type order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
}
type customer struct {
	id    int
	name  string
	email string
}

func (o order) printDetails() {
	fmt.Printf("Order ID: %d\n", o.id)
	fmt.Printf("Amount: %.2f\n", o.amount)
	fmt.Printf("Status: %s\n", o.status)
	fmt.Printf("Created At: %s\n", o.createdAt.Format(time.RFC1123))
}

// here also we can use pointer to change values ..

func newOrder(id int, amount float64, status string) order {
	return order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

func newCustomer(id int, name string, email string) customer {
	return customer{
		id:    id,
		name:  name,
		email: email,
	}
}

func (c customer) printDetails() {
	fmt.Printf("Customer ID: %d\n", c.id)
	fmt.Printf("Name: %s\n", c.name)
	fmt.Printf("Email: %s\n", c.email)
}
func main() {

	myOrder := newOrder(1, 100.50, "pending")
	myCustomer := newCustomer(1, "John Doe", "john.doe@example.com")

	// fmt.Printf("Order ID: %d\n", myOrder.id)
	// fmt.Printf("Amount: %.2f\n", myOrder.amount)
	// fmt.Printf("Status: %s\n", myOrder.status)
	// fmt.Printf("Created At: %s\n", myOrder.createdAt.Format(time.RFC1123))

	myOrder.printDetails()
	myCustomer.printDetails()

}
