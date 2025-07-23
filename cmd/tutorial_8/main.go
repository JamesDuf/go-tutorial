package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}    // Mutex to protect shared data since multiple goroutines will access it at around the same time
var wg = sync.WaitGroup{} // WaitGroup to wait for all goroutines to finish
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now() // start timer
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // increment the WaitGroup counter
		go dbCall(i) // start a goroutine to call the dbCall function
	}
	wg.Wait()                                              // wait for all goroutines to finish before proceeding
	fmt.Printf("\nTotal time taken: %v\n", time.Since(t0)) // print total time taken
	fmt.Printf("The results are: %v\n", results)           // print the results
}

func dbCall(i int) {
	// Simulate a database call with a random sleep time
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond) // simulate database call
	save(dbData[i])                                     // save the result of the db call
	log()                                               // log the results after each db call
	wg.Done()                                           // decrement the WaitGroup counter when the goroutine finishes
}

func save(result string) {
	m.Lock()                          // lock the mutex to protect shared data
	results = append(results, result) // append the result to the results slice
	m.Unlock()                        // unlock the mutex to allow other goroutines to access the shared data
}

func log() {
	m.RLock()                                            // lock the mutex for reading
	fmt.Printf("The current results are: %v\n", results) // print the results
	m.RUnlock()                                          // unlock the mutex after reading
}

//Multiple RLocks may be in effect at the same time, but only one Lock can be in effect at a time
//Lock waits for all full Locks to be released before acquiring the lock
