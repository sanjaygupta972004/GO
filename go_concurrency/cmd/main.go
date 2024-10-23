package main

import (
	"Go_Concurrency/internal/currency"
	"fmt"
	"sync"
	"time"
)

func runCurrencyWorker(
	workedId int,
	currencyChan <-chan currency.Currency,
	resultChan chan<- currency.Currency,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for c := range currencyChan {
		fmt.Printf("Worked %d started :\n", workedId)
		rates, err := currency.FetchCurrencyRates(c.Code)
		if err != nil {
			panic(err)
		}
		c.Rates = rates
		resultChan <- c
	}

}

func main() {
	ce := &currency.MyCurrencyExchange{
		Currencies: make(map[string]currency.Currency),
	}

	err := ce.FetchAllCurrencies()

	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}

	go func() {
		for {
			ce.Lock()
			usd, ok := ce.Currencies["usd"]
			ce.Unlock()
			if ok {
				fmt.Println("USD", usd.Rates)
			}
		}
	}()

	currencyChan := make(chan currency.Currency, len(ce.Currencies))
	resultChan := make(chan currency.Currency, len(ce.Currencies))

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go runCurrencyWorker(i, currencyChan, resultChan, &wg)
	}
	for _, curr := range ce.Currencies {
		currencyChan <- curr
	}

	close(currencyChan)

	startTime := time.Now()

	resultCount := 0

	for {
		if resultCount == len(ce.Currencies) {
			fmt.Println("Closing resultChan")
			break
		}
		select {
		case c := <-resultChan:
			ce.Lock()
			ce.Currencies[c.Code] = c
			ce.Unlock()

			resultCount++
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout")
			return
		}

	}

	// for code := range ce.Currencies {
	// 	wg.Add(1)
	// 	go func(code string) {
	// 		defer wg.Done()
	// 		rates, err := currency.FetchCurrencyRates(code)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		ce.Lock()
	// 		ce.Currencies[code] = currency.Currency{
	// 			Code:  code,
	// 			Name:  ce.Currencies[code].Name,
	// 			Rates: rates,
	// 		}
	// 		ce.Unlock()
	// 	}(code)

	// }

	endTime := time.Now()

	wg.Wait()

	fmt.Println("====result====")

	for _, curr := range ce.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}

	fmt.Println(startTime)

	fmt.Println("time taken :", endTime.Sub(startTime))

}
