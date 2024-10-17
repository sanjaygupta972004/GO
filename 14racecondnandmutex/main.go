package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println(" race condition and mutex")
	var result = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("first go routine")
		mut.Lock()
		result = append(result, 1)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("second go routine")
		mut.Lock()
		result = append(result, 3)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		wg.Done()
		fmt.Println("third go routine")
		//mut.Lock()
		result = append(result, 5)
		//mut.Unlock()
	}(wg, mut)

	wg.Wait()

	fmt.Println(result)

}
