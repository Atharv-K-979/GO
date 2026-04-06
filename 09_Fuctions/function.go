package main

import "fmt"


func add(a int,b int) int{
	return a+b
}

func getLanguages() (string, string, string){
	return "golang","python", "java"
}
 
func processIt(fn func(a int)int){
	fn(1)
}

func main(){
	fmt.Println(add(5, 3))
	fmt.Println(getLanguages())

	
	processIt(func(a int) int {
		fmt.Println("Processing", a)
		return a*2
	})
}