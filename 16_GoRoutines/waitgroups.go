package main

import (
	"fmt"
	"sync"
)

func task1(id int, w*sync.WaitGroup){
	defer w.Done()
	fmt.Printf("Task %d is running\n", id)
}
func main(){
	var wg sync.WaitGroup

	for i:=0; i<=10; i++{
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// WaitGroup  wg

// Think of wg like a counter:
// You increase it when starting a task
// You decrease it when a task finishes
// Program waits until it becomes 0

//     wg.Add(1)
// You’re telling Go:
// “One more goroutine is starting”
// So internally:
// counter = counter + 1

// defer wg.Done()
// When goroutine finishes:
// counter = counter - 1
// Using defer ensures it always runs at the end


// wg.Wait()
// Main function stops here until:
// counter == 0
// Only then program exits