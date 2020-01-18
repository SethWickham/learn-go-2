package main

import "fmt"

func directionalChannels() {
	fmt.Println("directionalChannels Print START")
	//bidirectional channel -by default a channel can send and recieve
	//but we can define a channel to only recieve or send see below
	c := make(chan string)

	go func() {
		c <- "channel to pass VALUES of TYPE string"
	}()
	//receive only channel
	cr := make(<-chan string)

	//send only channel
	cs := make(chan<- string)

	//set VALUE that has been passed to channel to ch and then print VALUE
	ch := <-c
	fmt.Println(ch)

	//TYPE print of c
	fmt.Printf("c channel TYPE: \t%T\n", c)
	//TYPE print of cr
	fmt.Printf("cr channel recieve TYPE: \t%T\n", cr)
	//TYPE print of cs
	fmt.Printf("cs channel send TYPE: \t%T\n", cs)

	fmt.Println("directionalChannels Print END")
}
