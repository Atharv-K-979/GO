package main

import (
	"fmt"
)

func main(){
	chan1 := make(chan int)
	chan2 := make(chan string)
	go func ()  {
		chan1 <- 42
	}()

	go func ()  {
		chan2 <- "Hello, World!"
	}()

	for i:=0; i<2; i++{
		select {
		case num := <- chan1:
			fmt.Println("Received number:", num)
		case msg := <- chan2:
			fmt.Println("Received message:", msg)
		}
	}
}