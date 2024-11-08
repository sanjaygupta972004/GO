package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func main() {
	fmt.Println("function in golang")
	result := Greeter()
	fmt.Println("value of greeter function", result)

	res := Adder(2, 5)

	fmt.Println("result of adder function ", res)

	total := ProAdder(2, 4, 6, 3, 5)

	fmt.Println("result of proFunction", total)

	fmt.Println("Here execute challenge functions")

	// number := []int{4, 5, 6, 8, 9, 4, 354, 334, 3}
	// maxNum := MaximumNumber(number)
	// minValue := MinimumValue(number...)
	// mean, _ := CalculateMean(number)

	// fmt.Printf("maximum number of array %d \n ", maxNum)
	// fmt.Printf("Minimum value of value %d\n", minValue)
	// fmt.Printf("Minimum value of value %f\n", mean)

	// amounts := []float64{300, 0, 1500, 30}

	// for _, amount := range amounts {
	// 	PaymentProcess(amount)
	// 	fmt.Println("Done Payment")
	// }

	data := PersonSlice{
		{Name: "Sanjay", Age: 20},
		{Name: "love", Age: 60},
		{Name: "Abhi", Age: 200},
		{Name: "Shivam", Age: 12},
	}
	fmt.Println("Execute function to check valid ")
	CheckValidAge(data)
	fmt.Println("Done")

}

func Greeter() string {
	return "Namaste golang "
}

func Adder(val1 int, val2 int) int {
	return val1 + val2
}

func ProAdder(value ...int) int {
	var total int
	for _, value := range value {
		total = total + value
	}
	return total

}
