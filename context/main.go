package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Learning context in golang")

	//let,s handle a scenario where we have to fetch some data from multiple api and handle all api cancelled if any api timeout or exceed
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	toCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	urls := []string{
		"https://github.com",
		"https://google.com",
		"https://bewanderic.com",
	}
	results := make(chan string)

	go toPerformTask()

	<-toCtx.Done()
	fmt.Println("tasked timeout")

	for _, url := range urls {
		go fetchApi(ctx, url, results)
	}

	for range urls {
		fmt.Println(<-results)
	}

}

func fetchApi(ctx context.Context, url string, results chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		results <- fmt.Sprintf("Error crating request for %s: %s", url, err)
		close(results)
	}

	client := http.DefaultClient

	resp, err := client.Do(req)

	if err != nil {
		results <- fmt.Sprintf("error getting during calling api %s :", err.Error())
		close(results)
	}

	defer resp.Body.Close()

	results <- fmt.Sprintf("Response from %s, %d", url, resp.StatusCode)
}

func toPerformTask() {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	}
}
