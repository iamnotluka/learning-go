package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	ninja := "Luka"

	smokeSignal := make(chan bool)
	
	go attack(ninja, smokeSignal)
	fmt.Println(<-smokeSignal)
}

func attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
	attacked <- true
}