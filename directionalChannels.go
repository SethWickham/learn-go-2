package main

import "fmt"

func directionalChannels() {
	fmt.Println("directionalChannels Print START")
	//basic channel
	c := make(chan string)

	go func() {
		c <- "channel to pass VALUES of TYPE string"
	}()
	//receiving channel
	cr := make(<-chan string)

	//send channel
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
