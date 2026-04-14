package main

import (
	"bufio"
	"os"
)

func main(){
	
	sourceFile, err := os.Open("./example.txt")  // to open a file
	if err != nil {
		panic(err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create("./example_copy.txt")  // to create a file
	if err != nil {
		panic(err)
	}
	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)  // to read from the source file
	writer := bufio.NewWriter(destFile)  // to write to the destination file

	for {
		b , err := reader.ReadByte()  // to read a byte from the source file
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err = writer.WriteByte(b)  // to write a byte to the destination file
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()  // to flush the writer buffer and write to the file
	println("File copied successfully.")

}