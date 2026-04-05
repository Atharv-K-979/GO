package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating map

	m := make(map[string]string)

	// setting an element
	m["name"] = "golang"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["name"], m["area"])
	// IMP: if key does not exists in the map then it returns zero value

	m1 := make(map[string]int)
	m1["age"] = 30
	m1["price"] = 50
	fmt.Println(m1["phone"])
	fmt.Println(len(m1))

	delete(m, "price")
	clear(m)  // delete all elements from the map

	fmt.Println(m)
	fmt.Println(m1)

	m2 := map[string]int{"price": 40, "phones": 3}
	print(m2)  
	v, ok := m["phones"]
	fmt.Println(v) // v gives value if not present then it gives zero value
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	m3 := map[string]int{"price": 40, "phones": 3}
	m4 := map[string]int{"price": 40, "phones": 8}
	fmt.Println(maps.Equal(m3, m4))

}