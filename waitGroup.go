package main

import (
	"fmt"
	"runtime"
)

func waitGroup() {
	//here we have code that is running in Parrallel thanks to the use of this computer's multiple core processors
	//go was created to take specific advantage of multiple core processor's and thus handles
	//Concurrency and Parrallelism extremely well
	fmt.Println("waitGroup Print START")

	fmt.Println("Number of CPU's:", runtime.NumCPU())

	fmt.Println("Goroutines", runtime.NumGoroutine())
	go numberToZero()
	go numberToTen()

	fmt.Println("Goroutines", runtime.NumGoroutine())

	fmt.Println("waitGroup Print END")
}

func numberToZero() {
	for i := 10; i >= 0; i-- {
		fmt.Println("numberToZero:", i)
	}

}

func numberToTen() {
	for i := 0; i <= 10; i++ {
		fmt.Println("numberToTen:", i)

	}
}
