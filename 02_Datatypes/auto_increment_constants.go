package main

import "fmt"

func main() {

	type Weekday int
	const (
		Sunday    Weekday = iota // 0
		Monday                   // 1
		Tuesday                  // 2
		Wednesday                // 3
		Thursday                 // 4
		Friday                   // 5
		Saturday                 // 6
	)

	// Bit flags with iota
	const (
		Read    = 1 << iota // 1  (001)
		Write               // 2  (010)
		Execute             // 4  (100)
	)
	// there thare is leftshift operator
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	fmt.Println(Read, Write, Execute)
	
}

// its like enum in C++ used to map things 