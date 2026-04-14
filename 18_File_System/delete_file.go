package main 

import (
	"os"
)

func main(){
	err := os.Remove("./example_copy.txt")  // to delete a file
	if err != nil {
		panic(err)
	}
	println("File deleted successfully.")
}
