package main

import "fmt"

func usingChannels() {
	fmt.Println("usingChannels Print START")
	c := make(chan string)

	//initializing a new go Routine
	go imSending(c)

	imReceiving(c)

	fmt.Println("usingChannels Print END")
}

// passing a VALUE to c
func imSending(c chan<- string) {
	c <- "Thirty-two"
}

// receiving the VALUE of c and print
func imReceiving(c <-chan string) {
	fmt.Println("The VALUE passed into channel:", <-c)
}
