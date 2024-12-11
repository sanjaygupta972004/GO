package main

import (
	"errors"
	"fmt"
	"time"
)

func printMessage(message string) error {
	if message == "" {
		return errors.New("message is not found")
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("index	 : %d Message :%s\n", i, message)
		time.Sleep(1 * time.Second)

	}

	return nil
}

func printNumber(num int) {
	fmt.Printf("goRoutine start here : %d\n", num)
	time.Sleep(1 * time.Second)
	fmt.Printf("goRoutine stop here : %d\n", num)

}

func main() {
	fmt.Println("studying about difference b/w concurrency and parallelism in golang")
	// run printMessage at concurrency situation
	// what i have observed here : without stop main goroutine we can't run other go routine

	// go printMessage("Hello concurrency")
	// time.Sleep(5 * time.Second)
	// printMessage("Hello parallelism")

	// run multiple go routine here
	// for i := 0; i < 5; i++ {
	// 	go printNumber(i)

	// }

	//channel()
	// Exa1()
	//Exa2()
	Exa3()
	time.Sleep(3 * time.Second)

}
