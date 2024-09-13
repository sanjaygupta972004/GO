package main

import (
	"fmt"
	"sort"
)

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


     myNumber := make([]int , 5  )

     myNumber[1] =45
     myNumber[0] = 34
     
    myNumber =  append(myNumber, 340 ,500, 9090)

   sort.Ints(myNumber)
   fmt.Println(myNumber)
}