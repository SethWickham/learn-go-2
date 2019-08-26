package main

import (
	"fmt"
	"runtime"
)

func waitGroup() {

	numberToZero()
	numberToTen()

	fmt.Println("Number of CPU's:", runtime.NumCPU())
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
