package main

import (
	"fmt"
	"sync"
)

var mut sync.Mutex

func counter(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	num := 0
	for i := 0; i < 50099900; i++ {
		mut.Lock()
		num++
		mut.Unlock()
	}
	fmt.Println("sending data....")
	ch <- num
	fmt.Println(" data sent ")
}

func Mutex() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go counter(ch, &wg)
	go counter(ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	// getting data from goroutine
	for data := range ch {
		fmt.Printf("Counter : value %d \n", data)
	}
}
