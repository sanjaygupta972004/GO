package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channel and deadlock")

	myChan := make(chan int, 3)
	wg := sync.WaitGroup{}

	wg.Add(2)
	// RECEIVED ONLY VALUE
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChan
		if isChannelOpen {
			fmt.Println(val)
		}
		fmt.Println(<-myChan)

		wg.Done()
	}(myChan, &wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {

		myChan <- 4
		myChan <- 9
		close(myChan)
		wg.Done()
	}(myChan, &wg)

	wg.Wait()
}
