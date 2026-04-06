// Variadic functions -- we can pass n no of arguments to a function 
package main

import "fmt"

func sum(nums ...int) int{  // veriadic parameter
	total := 0
	for _, n := range nums{
		total += n
	}
	return total
}

func main(){
	fmt.Println(sum(1, 2, 3)) // 6

	nums:= []int{4, 5, 6}
	fmt.Println(sum(nums...)) // 15 -- we need to unpack the slice using ... operator
}