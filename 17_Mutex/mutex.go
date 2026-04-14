package main

import (
	"fmt"
	"sync"
)

type post struct{
	views int 
	// to remove race
	mu sync.Mutex  // for read and write lock
}
func (p*post) inc(){
	p.views+=1;
}

func (p*post) inc1(waitGroup *sync.WaitGroup){
	// defer waitGroup.Done()
	defer func(){
		p.mu.Unlock()
		waitGroup.Done()
	}()
	p.mu.Lock() // lock the mutex before modifying the shared data
	p.views+=1;


	// writing unlock in the fun is good as if process ends in betn then it wont be unlock 
	// p.mu.Unlock()   // unlock the mutex after modifying the shared data	
}


func main(){
	myPost:=post{views:0}

	// myPost.inc()
	// myPost.inc()
	// myPost.inc()

	var waitGroup sync.WaitGroup
	for i:=0; i<100;i++{
		waitGroup.Add(1)
		go myPost.inc1(&waitGroup)   // race condition
	}
	waitGroup.Wait()
	fmt.Println(myPost.views)

}

/// Race condition is a situation where multiple threads or goroutines access and modify shared data concurrently, leading to unpredictable results. 
// In the above code, if we were to run the inc() method concurrently without proper synchronization, we could end up with