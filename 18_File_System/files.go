package main

import (
	"os"
)

func main(){
	f, err := os.Open("./example.txt")  // to open a file
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()  // to get file info
	if err != nil {
		panic(err)
	}
	println("File Name:", fileInfo.Name())
	println("File Size:", fileInfo.Size())
	println("File Mode:", fileInfo.Mode())
	// println("File ModTime:", fileInfo.ModTime())
	println("Is Directory:", fileInfo.IsDir())

	// to read a file
	buf := make([]byte, fileInfo.Size())
	_, err = f.Read(buf)
	if err != nil {
		panic(err)
	}
	println("File Content:", string(buf))

	// shortcut but not for big files
	content, err := os.ReadFile("./example.txt")
	if err != nil {
		panic(err)
	}
	println("File Content:", string(content))

	defer f.Close()  // to close the file after reading


	// read folder
	dir, err := os.Open("../")  // to open a directory
	if err != nil {
		panic(err)
	}
	files, err := dir.Readdir(-1)  // to read all files in the directory  -1 means all files 1 means 1 type of file
	if err != nil {
		panic(err)	
	}
	for _, file := range files {
		println("File Name:", file.Name())
		println("File Size:", file.Size())
		println("File Mode:", file.Mode())
		println("Is Directory:", file.IsDir())
		println("-----")
	}					
	
	
	
}