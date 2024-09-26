package main

import (
	"fmt"
	"sort"
)
// veriadic concept

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
	    total += number
	}
	return total
    }

func main() {
	fmt.Println("welcome to slices")
	var fruitList =  []string{
		"mango", "orange",
	}
     
	fmt.Printf("data type is : %T\n" , fruitList)

	fmt.Println("list of fruit", fruitList)

	fruitList = append(fruitList, "peach", "banana")

	fmt.Println(fruitList)

	  midValues := fruitList[1:3]

	  fmt.Println("mid value of slice :", midValues)

      numberSlice := []int {
	1,12,3,44,5,06,27,8,
      }


     lengthSlice := len(numberSlice)

     fmt.Println(lengthSlice)


     myNumber := make([]int ,5 , 20 )

     myNumber[1] =45
     myNumber[0] = 34
     
    myNumber =  append(myNumber, 340 ,500, 9090)

   sort.Ints(myNumber)
   fmt.Println(myNumber)
  

   //  how to delete an element using index method
   course := []string{"reactJs", "typeScript","ruby","rust", "go"}

   fmt.Println(course)

   var index int = 3

   course = append(course[:index] , course[index+1])

   fmt.Println("new course",course)

}