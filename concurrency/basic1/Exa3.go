package main

import (
	"fmt"
	"sync"
)

func sumChunks(sum []int, totalChunk chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine finishes
	total := 0

	for _, val := range sum {
		total += val
	}
	totalChunk <- total
}

func Exa3() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	chunksGoRoutine := 3
	ch := make(chan int, chunksGoRoutine)

	sizeChunks := len(array) / chunksGoRoutine
	if len(array)%chunksGoRoutine != 0 {
		sizeChunks++
	}

	var wg sync.WaitGroup
	for i := 0; i < chunksGoRoutine; i++ {
		start := i * sizeChunks
		end := (i + 1) * sizeChunks

		if end > len(array) {
			end = len(array)
		}

		wg.Add(1) // Increment the counter for each goroutine
		go sumChunks(array[start:end], ch, &wg)
	}

	// Close the channel once all goroutines have finished
	go func() {
		wg.Wait() // Wait for all goroutines to finish
		close(ch) // Close the channel
	}()

	totalResult := 0
	for val := range ch {
		data := val
		fmt.Println("value from goroutine")
		totalResult += data
	}
	fmt.Printf("end of code and printing result %d \n", totalResult)
}
