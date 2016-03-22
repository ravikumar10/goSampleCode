
// This sample program demonstrates how to create goroutines
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main is the entry point for all Go programs.
func main() {
	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)
	// Allocate a context for every available core.
	//runtime.GOMAXPROCS(runtime.NumCPU())
	// Allocate two contexts for the scheduler to use.
	runtime.GOMAXPROCS(2)
	fmt.Println("Start Goroutines")

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		for count := 1; count <= 1000; count++ {
			fmt.Printf("A%d ", count)
		}
	}()

	// Declare an anonymous function and create a goroutine.
	go func() {
		// Schedule the call to Done to tell main we are done.
		defer wg.Done()
		for count := 1000; count >= 1; count-- {
			fmt.Printf("B%d ", count)
		}
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
