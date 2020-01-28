// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to show how to create race conditions in
// our programs. We don't want to do this.
package main

import (
	"fmt"
	"sync"
)

// counter is a variable incremented by all goroutines.
var counter int32

func main() {

	// Number of goroutines to use.
	const grs = 2

	// wg is used to manage concurrency.
	var wg sync.WaitGroup
	wg.Add(grs)
	var mu sync.Mutex

	// Create two goroutines.
	for g := 0; g < grs; g++ {
		go func() {

			//mutex example
			for i := 0; i < 2; i++ {
				mu.Lock()
				{
					// Capture the value of Counter.
					value := counter
					// Increment our local value of Counter.
					value++
					fmt.Println("tracing") //this is extra latency that we don't need! don't log in the mutex, get in and out as fast as possible
					// Store the value back into Counter.
					counter = value
				}
				mu.Unlock()
			}

			// // atomic thing
			// for i := 0; i < 2; i++ {
			// 	atomic.AddInt32(&counter, 1)
			// 	fmt.Println("tracing") //race condition trigger because of context switching
			// }

			//BAD! race condition
			// for i := 0; i < 2; i++ {
			// 	// Capture the value of Counter.
			// 	//value := counter
			// 	// Increment our local value of Counter.
			// 	//value++
			// 	fmt.Println("tracing") //race condition trigger because of context switching
			// 	// Store the value back into Counter.
			// 	//counter = value
			// }

			wg.Done()
		}()
	}

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

/*
 go build -race
 ./race-condition

==================
WARNING: DATA RACE
Read at 0x0000011a5118 by goroutine 7:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:33 +0x4e
Previous write at 0x0000011a5118 by goroutine 6:
  main.main.func1()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:39 +0x6d
Goroutine 7 (running) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3
Goroutine 6 (finished) created at:
  main.main()
      /Users/bill/code/go/src/github.com/ardanlabs/gotraining/topics/go/concurrency/data_race/example1/example1.go:43 +0xc3
==================
Final Counter: 4
Found 1 data race(s)
*/
