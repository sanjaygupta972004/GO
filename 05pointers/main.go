package main

import "fmt"

func main() {
	fmt.Println("learning about pointer in golang")

	var prt *int

	fmt.Println("value of pointer", prt)

	myNumber := 25

	prt = &myNumber
         
	fmt.Println("reference of memory  predefined variable as pointer", *prt)
	ptr := *prt +4


	fmt.Println("reference of memory  predefined variable as pointer", prt)
	fmt.Println("value of pointer variable ", *prt)
	fmt.Println(ptr)
         
         var pointer * bool
	isLagged := true

	pointer = &isLagged

	fmt.Println("reference of memory  predefined variable as pointer",pointer )
	fmt.Println("value of pointer variable ", *pointer) 
}