package main

import "fmt"

func main() {
	fmt.Println("learning loop in go")
	myDays := []string{"Sunday", "Tuesday", "Wednesday", "Thursday"}
	fmt.Println(myDays)
	// first method to implement loop
	for _, value := range myDays {
		fmt.Println("Value of days", value)
	}

	for d := 0; d < len(myDays); d++ {
		fmt.Println(myDays[d])
	}

	initialValue := 1

	for initialValue < 10 {

		if initialValue == 4 {
			initialValue++
			continue

		}
		initialValue++

		fmt.Println(initialValue)

	}
}
