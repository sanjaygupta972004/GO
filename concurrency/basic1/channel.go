package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int, value int) {
	ch <- value
}

// send multiple data
func sendMultiData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(2 * time.Second)
	}
	close(ch)
}
func channel() {
	ch := make(chan int)
	go sendData(ch, 30)
	go sendData(ch, 39)

	fmt.Println("Received data from first go routine", <-ch)
	fmt.Println("Received data from second  go routine", <-ch)

	// get multiple data from chan
	go sendMultiData(ch)
	for val := range ch {
		fmt.Printf("Received multi data from channel %d\n", val)
	}

}
