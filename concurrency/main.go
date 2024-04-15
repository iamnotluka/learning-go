package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	ninjas := []string{"Luka", "Relja", "Tamara", "Aleks"}

	for _, ninja := range ninjas {
		go attack(ninja)
	}

	time.Sleep(time.Second)
}

func attack(target string) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
}