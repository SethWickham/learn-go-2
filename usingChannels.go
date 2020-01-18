package main

import "fmt"

func usingChannels() {
	fmt.Println("usingChannels Print START")
	c := make(chan int)

	//initializing a new go Routine
	go imSending(c)

	imReceiving(c)

	fmt.Println("usingChannels Print END")
}

// passing a VALUE to c
func imSending(c chan<- int) {
	c <- 32
}

// receiving the VALUE of c and print
func imReceiving(c <-chan int) {
	fmt.Println("The VALUE passed into channel:", <-c)
}
