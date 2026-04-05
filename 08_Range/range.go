package main

import "fmt"

func main(){
	// range -> gives index and value of each element in the collection
	nums := []int{10, 20, 30}
	for i:=0 ; i<len(nums); i++{
		fmt.Println(i, nums[i])
	}

	//with range
	for i, v := range nums{
		fmt.Println(i, v)
	}
	
	// if you dont want index
	for _, v := range nums{
		fmt.Println(v)
	}

	m := map[string]int{"price": 40, "phones": 3}
	for k, v := range m{
		fmt.Println(k, v)
	}

	for _,c:= range "Atharv"{
		fmt.Println(c)  // gives unicode value of the character
	}
	for _, c := range "Atharv"{
		fmt.Printf("%c\n", c)  // gives character
	}
}