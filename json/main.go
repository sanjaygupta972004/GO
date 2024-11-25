package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Student struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Section    string `json:"section"`
	IsFeesPaid bool   `json:"is_fees_paid"`
}

func main() {
	log.Println("working with json package in golang")

	std := Student{Name: "sunny", Address: "Dulapur", Section: "A", IsFeesPaid: true}

	fmt.Println("logging student data", std)

	// how to convert plain text object into json data

	jsonData, err := json.Marshal(std)
	if err != nil {
		fmt.Println("error marshalling data", err)
	}

	fmt.Println("logging jsonData", string(jsonData))

	// how to convert json data into plain text
	var data Student
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(" error logging plain data after converting from json", err)
	}

	fmt.Println("logging plain data after converting from json", data)

}
