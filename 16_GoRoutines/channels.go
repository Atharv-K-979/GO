package main

import (
	"fmt"
)

func processNum(numChan chan int) {
	num := <-numChan // receive number from channel
	fmt.Println("Received number:", num)
}

func processNum1(numChan chan int) {
	for {
		num := <-numChan // receive number from channel
		fmt.Println("Received number:", num)
	}
}

// func sum(numresult chan int, a, b int) {
// 	numresult <- a + b
// 	result := <-numresult
// }
func sum(numresult chan int, a, b int) {
	numresult <- a + b
}


func task(done chan bool) {
	// fmt.Println("Task is running")
	// done <- true // signal that task is done

	defer func() {
		done <- true // signal that task is done
	}()

	fmt.Println("Task is running")
}
func main() {
	//***************** 1 *****************

	// messChannel := make(chan string)

	// messChannel <- "ping"

	// msg := <- messChannel

	// fmt.Println(msg)
	//error: all goroutines are asleep - deadlock!

	//***************** 2 *****************
	// numChan := make(chan int)

	// go processNum(numChan)
	// // numChan <- 42

	// time.Sleep(time.Second*2)

	// for{
	// 	numChan <- rand.Intn(100)
	// }

	//***************** 3 *****************
	result := make(chan int)
	go sum(result, 4, 5)

	res := <-result
	fmt.Println("Sum is:", res)

	//***************** 4 *****************
	done := make(chan bool)
	go task(done)

	<-done // wait for task to signal that it's done
	fmt.Println("Task completed")
}

//
// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
// 	/home/atharv/Documents/GO/16_GoRoutines/channels.go:10 +0x36
// exit status 2
