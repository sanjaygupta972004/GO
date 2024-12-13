package main

import (
	"fmt"
	"sync"
)

func sender(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		fmt.Printf("sending data %d\n", i)
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed, exiting receiver.")
			break
		}

		fmt.Printf("result : %d\n", value)
	}
}

func exa1() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go sender(ch, &wg)
	receiver(ch)
	defer func() {
		wg.Wait()
	}()
}
