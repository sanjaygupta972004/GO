package main

import "fmt"


func main() {
	fmt.Println("function in golang")
       result := Greeter()
       fmt.Println("value of greeter function", result)

       res := Adder(2,5)

       fmt.Println("result of adder function ", res)

       total := ProAdder(2 ,4,6,3,5)

       fmt.Println("result of proFunction",total)

}


func Greeter () string {
	return "Namaste golang "
}

func Adder( val1 int , val2 int) int {
	return val1+val2
}


func ProAdder(value ...int)  int {
	var total int 
	for _ , value := range value {
		total = total+value
	}
	return  total

} 
