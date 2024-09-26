package main

import "fmt"


func main() {
 fmt.Println("Learning map in golang")

 // how to define map

 //var myMap map[string]int 

 myLang := make(map[string]string)

 myLang["js"] = "javaScript"
 myLang["ts"] = "typeScript"
 myLang["g"] = "go"

 fmt.Println("the value of myLang",myLang)

 // how to delete a value from map

 delete(myLang, "g")

 fmt.Println("the value of myLang",myLang)

 
 for key , value := range myLang {
	fmt.Printf("this is my : %v language %v", key ,value)
 }




}