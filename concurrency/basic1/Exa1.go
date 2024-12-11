package main

import "fmt"

// create a function where a counter update number and read in main file using goroutine
func counter(ch chan int) {
	num := 0
	for i := 0; i < 5; i++ {
		if i < 5 {
			num++
			ch <- num
		}
	}

	close(ch)
}

func Exa1() {
	ch := make(chan int)

	go counter(ch)

	for val := range ch {
		fmt.Printf("Getting value from counter %d\n ", val)
	}

	fmt.Println("end example first function ")
}
