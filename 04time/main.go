package main

import (
	"fmt"

	"time"	
)



func main(){
	fmt.Println("learning about time of golang")
        presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 monday"))

	createdDate := time.Date(2000, time.August, 10, 32 ,40 ,0 ,0, time.UTC)

	fmt.Println(createdDate)
        
}