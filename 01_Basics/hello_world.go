package main  // every go program has belong to  main means the executable program
import "fmt" // fmt imports the formatting package used for Println() Print() Printf()  

func main() { // Entry point of the program
    fmt.Println("Hello, World!")          // print then next line
	fmt.Print("Atharv")                   // print only

	fmt.Printf("Integer: %d\n",5);        // print + we can use format specifiers 
	fmt.Printf("Character:  %c\n",'a')
	fmt.Printf("Float:  %f\n",5.7)
	fmt.Printf("String:  %s\n","String")

	// age:=19;
	// age="Arhav";
	// fmt.Print(age);            not allowed


	// tmp:="Athav";
	// tmp=19;
	// fmt.Print(tmp);            not allowed
}



// Run the file 

// go run hello_world.go

// or

// go build hello_world.go
// ./hello

