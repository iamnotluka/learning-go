package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Luka")
	fmt.Println(message)
}