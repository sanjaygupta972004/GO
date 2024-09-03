package main

import (
	"fmt"
)

const  JwtToken string = "bob pratap singh"


func main(){
  fmt.Println("variables")	
   var username string = " sanjay gupta"
   fmt.Printf("type of variable : %T \n", username)
   
   var isLogin bool = false
   fmt.Printf("type of variable : %T \n", isLogin)

   var noOfUser uint = 8593583958353
 
   fmt.Printf("Variable is of type: %T \n", noOfUser)
   fmt.Println(noOfUser)
    
   var isuser int 
   fmt.Printf("Variable is of type: %T \n", isuser)

   var logout = true
   fmt.Printf("Variable is of type: %T \n", logout)

   dataBaseConnetionString := "xccjzzsdssckzjdjaifd"
   dataBaseConnetionString = "idfuis"

   fmt.Printf("Variable is of type: %T \n", dataBaseConnetionString)
   fmt.Println(dataBaseConnetionString)
 
   fmt.Printf("Variable is of type: %T \n", JwtToken)
   fmt.Println(JwtToken)
} 