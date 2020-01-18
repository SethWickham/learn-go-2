package main

import "fmt"

// channels allow us to communicate between our go routines, they are the only tool we have in go to communicate
// between different go routines, we can send our data into a channel and thereby use that data in any go routine that is connected to our channel
// Like in the rest of go once a Channel's TYPE is defined we cannot pass in VALUES of different TYPES into it, and a channel is itself a TYPE in go

//  VERY IMPORTANT- Channels block, they must, typically it is good practice to build our code in a way that is unbuffered
// so that there is always the interlocking component where a pass will happen - Todd Mcleod
// in order to pass VALUES between GO ROUTINES we use CHANNELS
func channels() {
	fmt.Println("channels Print START")

	//here we create our channel
	ch := make(chan int)
	// without the go func our code will not run and will get blocked
	go func() {

		// placing our value onto a channel
		ch <- 32

	}()
	// taking the value off of the channel
	fmt.Println(<-ch)
	fmt.Printf("%T\n", ch)

	fmt.Println("channels Print END")
}
