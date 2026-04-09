package main

//enums - enumerated types - a way to define a set of named constants

import "fmt"



// func orderStatus(status string) string {
// 	fmt.Println("Changing order status ", status)
// 	return status
// }

//enum types
type orderStatus int
const (
	pending   orderStatus = iota    // default value is 0
	confirmed // 1
	shipped   // 2
	delivered // 3
	canceled  // 4
)

// using string instead of int
// type orderStatus string
// const (
// 	pending   orderStatus = "pending"
// 	confirmed orderStatus = "confirmed"
// 	shipped   orderStatus = "shipped"
// 	delivered orderStatus = "delivered"
// 	canceled  orderStatus = "canceled"
// )


func changeOrderStatus(status orderStatus) {
	fmt.Println("Changing order status ", status)
}
func main() {
	changeOrderStatus(confirmed)  
	changeOrderStatus(canceled)
}
