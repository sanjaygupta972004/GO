package main

import (
	"fmt"
	"go-rabbitmq/receive"
	"go-rabbitmq/send"
)

func main() {
	fmt.Println("Working with rabbitmq in golang")

	send.Producer()
	receive.Consumer()

}
