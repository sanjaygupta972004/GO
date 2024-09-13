package main

import (
	"fmt"

)

func main() {
	fmt.Println("started learning Array in go")
	var fruitList [3]string

	fruitList[0] = "mango"
	fruitList[1] = "banana"
	fruitList[2] = "grapes"

	fmt.Println("List of fruitList is :", fruitList )
	fmt.Println("length of fruit: " , len(fruitList))
	

	var numberList = [5]int{1, 2, 3, 5}

        fmt.Println("list of number : ",numberList)

	for index , number :=  range numberList {
		fmt.Println(index , number)
	}				


}
