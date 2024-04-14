package main

import "fmt"

var points = []int{20, 90, 100, 45, 70}

var score = 99.5

func sayHello(n string) {
	fmt.Println("hello", n)
}

func showScore() {
	fmt.Println("You've scored this many points:", score)
}