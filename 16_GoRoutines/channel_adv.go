package main

import (
	"fmt"
	"time"
)


func emailSender(emailChan chan string, done chan bool) {
	defer func() {
		done <- true // signal that email sender is done
	}()
	for email:= range emailChan {
		fmt.Println("Sending email to:", email)
		time.Sleep(time.Second)
	}
}

func main(){
	emailChan := make(chan string,100)
	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// for more emails
	done:= make(chan bool)
	go emailSender(emailChan, done)
	for i:=0 ; i<100 ; i++{
		emailChan <- fmt.Sprintf("%d@example.com", i+1)
	}
	close(emailChan) // close channel to signal no more emails  this is imp
	<-done
	
}