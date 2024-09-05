package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("welcome to bewanderic")
	fmt.Println("pls rete travelPackage between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for rating ", input)

	newRating,err :=  strconv.ParseFloat(strings.TrimSpace(input) , 64)

	if err != nil {
            fmt.Println("getting error", err.Error())
	}else {
		fmt.Println("add 1 in rating", newRating+1)
	}



}