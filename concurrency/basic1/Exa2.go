package main

import "fmt"

func ping(ch chan string, msg string) {
	if msg == "" {
		fmt.Println("Message is not found")
	}
	fmt.Println("Sending msg...")
	ch <- msg
	fmt.Println("Message sent")
	close(ch)
}

func pong(ch chan string) {
	fmt.Println("Receiving msg from channel")
	data := <-ch
	fmt.Println("Received msg ", data)
}

func Exa2() {
	ch := make(chan string)

	go pong(ch)
	go ping(ch, "Hello go developer, what's up !")
}
