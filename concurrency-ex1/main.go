package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d completed\n", id)
}

func waitGroupExample() {
	fmt.Println("\n-waitGroupExample:")
	var wg sync.WaitGroup

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // Increment the counter before starting each goroutine
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish

	fmt.Println("-waitGroupExample: All workers have completed.")
	fmt.Println()
}

// ------ mutex

func mutexExample() {
	fmt.Println("\n-mutexExample:")
	var (
		counter int
		mutex   sync.Mutex
	)

	var wg sync.WaitGroup

	numGoroutines := 3
	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()         // Acquire the lock
			defer mutex.Unlock() // Ensure the lock is released when the function exits
			counter++
		}()
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Printf("-mutexExample: Final Counter Value: %d\n\n", counter)
}

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	waitGroupExample()
	mutexExample()
}
