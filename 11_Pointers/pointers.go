package main

import "fmt"

func changeNum(num int){ // by value passing
	num = 10
	fmt.Println("Inside changeNum:", num) // 10
}

func changeNum1(num *int){ // by reference passing
	*num = 10                                          //  * dereferencing pointer to change value at that address
	fmt.Println("Inside changeNum:", *num) // 10
}
func main(){
	// value
	// num := 5
	// fmt.Println("Before changeNum:", num) // 5
	// changeNum(num)
	// fmt.Println("After changeNum:", num) // 5

	// reference
	num := 5
	fmt.Println("Before changeNum:", num) // 5
	changeNum1(&num)
	fmt.Println("After changeNum:", num) // 10

}