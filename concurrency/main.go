package main

import (
	"fmt"
	"time"
)

func main() {
	ninjas := []string{"Luka", "Relja", "Tamara", "Aleks"}

	for _, ninja := range ninjas {
		attack(ninja)
	}
}

func attack(target string) {
	fmt.Println("Throwing ninja starts at", target)
	time.Sleep(time.Second)
}