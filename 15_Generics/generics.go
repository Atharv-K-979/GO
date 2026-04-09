package main

import "fmt"

// func printSlice(items[] int){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }
// func printStringSlice(items[] string){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// generics - a way to write functions and data structures that can work with any type

func printSlice[T any](items[] T){  // in go 1.18 we can use any type with generics    any or interface{} 
	for _, item := range items {
		fmt.Println(item)
	}
}

// can be used on struct
// type stack struct{
// 	items []int
// }

// type stack1 struct{
// 	items []string
// }

//with generics
type stack[T any] struct{
	items []T
}

// also comporables can be used with generics 


func main() {
	// intSlice := []int{1, 2, 3, 4, 5}
	// printSlice(intSlice)
	// // string slices
	// stringSlice := []string{"a", "b", "c", "d", "e"}
	// printStringSlice(stringSlice)

	// now if if i want boolean then i have to create new fun so duplication is there

	// with generics
	intSlice := []int{1, 2, 3, 4, 5}
	printSlice(intSlice)
	// string slices
	stringSlice := []string{"a", "b", "c", "d", "e"}
	printSlice(stringSlice)
	// boolean slices
	boolSlice := []bool{true, false, true, false}
	printSlice(boolSlice)

	// stack of int
	intStack := stack[int]{}
	intStack.items = []int{1, 2, 3, 4, 5}
	fmt.Println(intStack)

	// stack of string
	stringStack := stack[string]{}
	stringStack.items = []string{"a", "b", "c", "d", "e"}
	fmt.Println(stringStack)
}