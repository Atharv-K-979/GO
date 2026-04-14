package main

import (
	"os"
)

func main(){
	f, err := os.Create("./example2.txt")  // to create a file
	if err != nil {
		panic(err)
	}
	defer f.Close()  // to close the file after writing

	content := "Hello, this is an example file. "
	_, err = f.WriteString(content)  // to write to the file
	if err != nil {
		panic(err)
	}

	bytes:= []byte("Atharv Kulkarni")
	_, err = f.Write(bytes)  // to write bytes to the file
	if err != nil {
		panic(err)
	}
	println("File created and content written successfully.")
}