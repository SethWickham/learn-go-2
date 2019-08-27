package main

import (
	"fmt"
	"runtime"
	"sync"
)

// The below code explores the use of writing Concurrent and Parrallel code in go
// this allows us to fully utilize the computer's multiple core Processors as we see fit
// we are using the package runtime to count our Goroutines and CPU's
// and we are using package sync to access WaitGroup which allows us to cause our
// Goroutines to wait until certain actions are completed, thus allowing us to sync up our program
// so that it behaves as we desire

//it is important to note that code may be concurrent and not running Parrallel if the computer only has a single processor

var wg sync.WaitGroup

func waitGroup() {
	//here we have code that is running in Parrallel thanks to the use of this computer's multiple core processors
	//go was created to take specific advantage of multiple core processor's and thus handles
	//Concurrency and Parrallelism extremely well
	fmt.Println("waitGroup Print START")
	//this will show us how many processor's are available to us on our computer
	fmt.Println("Number of CPU's:", runtime.NumCPU())

	fmt.Println("starting number of Goroutines:", runtime.NumGoroutine())
	//here we are launching new go routines (to take advantage of multiple core processors ) by using the word go before our functions
	// we use wg.Add to add our GoRoutines to our WaitGroup
	wg.Add(2)
	go numberToZero()
	go numberToTen()
	fmt.Println("new number of Goroutines:", runtime.NumGoroutine())
	wg.Wait()

	//because our newly added (2) Goroutines have ended we are back to one Goroutine which runs our func main, then when main exits our program shuts down
	fmt.Println("Final number of Goroutines:", runtime.NumGoroutine())

	fmt.Println("waitGroup Print END")
}

// wg.Done is basically saying thanks for waiting you can stop waiting now and exit the WaitGroup.
func numberToZero() {
	for i := 10; i >= 0; i-- {
		fmt.Println("numberToZero:", i)
	}
	wg.Done()
}

func numberToTen() {
	for i := 0; i <= 10; i++ {
		fmt.Println("numberToTen:", i)
	}
	wg.Done()
}
