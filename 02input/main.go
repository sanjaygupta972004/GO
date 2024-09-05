package main

import (
	"bufio"
	"fmt"
	"os"
)


func sum(a,b int)(int){
        res := a+b 
	return res
}

func main() {
	fmt.Println(" learning input")

	var welcome string =  " welcome to input user"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" enter the rating for package :")
	//  comma ok // error syntax coming 
	input,err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("something wrong when reading input")
		return 
	}
	fmt.Println("thanks for input",input)
	fmt.Printf(" types of input variable : %T \n", input)

	// call outer function in main function
	 res := sum(4,6)
	 fmt.Println("the sum of two number", res)

}

