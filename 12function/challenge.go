package main

import (
	"errors"
	"log"
)

func MaximumNumber(number []int) int {
	if len(number) == 0 && number == nil {
		return 0
	}
	max := number[0]
	for _, num := range number {
		if max < num {
			max = num
		}
	}
	return max
}

func MinimumValue(number ...int) int {
	min := number[0]

	for _, num := range number {
		if min > num {
			min = num
		}
	}
	return min
}

func CalculateMean(number []int) (float32, error) {
	if number == nil {
		return 0, errors.New("number is not valid")
	}

	sum := 0

	for _, num := range number {
		sum += num
	}

	mean := float32((sum) / len(number))

	return mean, nil
}

// how to handle error in golang : golang also famous to handling error it provides some inbuilt function like panic defer and recover using this we can easily handle error

func PaymentProcess(amount float64) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Some thing went to wrong while processing payment", r)
		}
	}()
	// handle condition when amount is 0
	if amount == 0 {
		panic("can't pay 0 rupees")
	}
	// handle network error
	if amount > 1300 {
		panic("Network error")
	}

	// process payment logic
	log.Printf("This is amount can pay %f", amount)
}

// solving a problem on handing defer and recover and panic

func CheckValidAge(person []Person) {
	defer func() {
		r := recover()
		if r != nil {
			log.Printf("Checking it valid error : %s\n", r)
		}
	}()

	for _, per := range person {
		if per.Age > 100 {
			panic("Operation Panicked")
		} else {
			log.Printf("Operational successful")
		}

	}
}
