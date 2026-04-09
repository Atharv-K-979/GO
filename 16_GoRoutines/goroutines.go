package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Printf("Task %d is running\n", id)
}

func main(){
	for i:=0 ; i<10; i++{
		// task(i)  // this will run sequentially
		go task(i) // jab tak run hua as it's sooo fast we come outside main and exit so no output


		go func(i int){
			fmt.Println(i)
		}(i)
	}
	time.Sleep(1 * time.Second) // to wait for all goroutines to finish before exiting main
	// so order is not guaranteed as they are running concurrently

}