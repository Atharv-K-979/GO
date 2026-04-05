package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("person is an adult")
	} else {
		fmt.Println("person is not an adult")
	}
	// age := 10

	if age >= 18 {
		fmt.Println("person is an adult")
	} else if age >= 12 {
		fmt.Println("person is teenager")
	} else {
		fmt.Println("person is a kid")
	}

	var role = "admin"
	var hasPermissions = false

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	// go does not have ternary, you will have to use normal if else

}

func doSomething() (any, any) {
	panic("unimplemented")
}