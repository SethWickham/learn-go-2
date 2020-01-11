package main

import "fmt"

//  VERY IMPORTANT- Channels block, they must, typically it is good practice to build our code in a way that is unbuffered
// so that there is always the interlocking component where a pass will happen - Todd Mcleod
// in order to pass VALUES between GO ROUTINES we use CHANNELS
func channels() {
	fmt.Println("channels Print START")

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
