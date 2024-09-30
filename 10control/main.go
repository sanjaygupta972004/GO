package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switch case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of diceNumber:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("now you can open your spot")
		fallthrough

	case 2:
		fmt.Println("this is your primary number")

	case 3:
		fmt.Println("this is mid number of dis")
		default :
		fmt.Println("what happen with you !")
	}

}
