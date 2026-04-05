package main

import "fmt"

// for -> only construct in go for looping
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	println("1")
	// }

	// classic for loop
	for i := 0; i <= 3; i++ {
		// break
		if i == 2 {
			continue
		}
		fmt.Println(i)

	}

	// version 1.22 range
	for i := range 11 {
		fmt.Println(i)   // 0 to 2
	}
	
	// for i:= range "Atharv"{
	// 	fmt.Println(i);
	// 	fmt.Println("hii")
	// }
	// wrong
	for _, ch := range "Atharv" {  // _ is ignored
    fmt.Printf("%c\n", ch)
}
}