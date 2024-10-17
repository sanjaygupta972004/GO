package main

import (
	"fmt"
	"net/http"
	"sync"
	//"time"
)

//	func greeter(s string) {
//		for i := 0; i < 4; i++ {
//			time.Sleep(5 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
var wg sync.WaitGroup // usually this is pointer

var mu sync.Mutex

var signal []string

func getStatusCode(endPointes string) {
	defer wg.Done() // this flag responsible to decrease go routine

	res, err := http.Get(endPointes)

	if err != nil {
		fmt.Println("some thing went wrong", err.Error())
	} else {
		mu.Lock()
		signal = append(signal, "this is shared memory")
		mu.Unlock()
		fmt.Printf(" %d status code this %s \n", res.StatusCode, endPointes)
	}
}
func main() {
	fmt.Println("go routine")

	webLists := []string{
		"http://bewanderic.com",
		"http://github.co",
		"http://google.com",
		"http://chaicode.com",
		"http://facebook.com",
	}

	for _, web := range webLists {
		wg.Add(1) // this is responsible for adding go routine
		getStatusCode(web)

	}

	wg.Wait() // this is responsible for execute all go routine

	// go greeter("hello")
	// greeter("world")
}
