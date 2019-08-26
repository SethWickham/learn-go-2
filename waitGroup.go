package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func waitGroup() {
	//here we have code that is running in Parrallel thanks to the use of this computer's multiple core processors
	//go was created to take specific advantage of multiple core processor's and thus handles
	//Concurrency and Parrallelism extremely well
	fmt.Println("waitGroup Print START")
	//this will show us how many processor's are available to us on our computer
	fmt.Println("Number of CPU's:", runtime.NumCPU())

	fmt.Println("Goroutines:", runtime.NumGoroutine())
	//here we are launching new go routines (to take advantage of multiple core processors ) by using the word go before our functions
	wg.Add(2)
	go numberToZero()
	go numberToTen()
	fmt.Println("new number of Goroutines:", runtime.NumGoroutine())
	wg.Wait()

	//because our newly added (2) Goroutines have ended we are back to our one Goroutine which runs our func main, then when main exits our program shuts down
	fmt.Println("Final number of Goroutines:", runtime.NumGoroutine())

	fmt.Println("waitGroup Print END")
}

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
