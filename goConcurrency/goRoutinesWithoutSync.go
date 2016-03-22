

//goroutines without synchronization
//This is a wrong implementation
package main

import (
	"fmt"
)

// main is the entry point for all Go programs.
func main() {
	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {

		for count := 1; count <= 100; count++ {
				fmt.Printf("%d ", count)
		}
	}()

	go func() {

		for count := 100; count >= 1; count-- {
			fmt.Printf("%d ", count)
		}
	}()

	fmt.Println("\nTerminating Program")
}
